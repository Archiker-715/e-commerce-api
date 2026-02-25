package uc

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	ctxpkg "github.com/Archiker-715/e-commerce-api/internal/auth/ctx"
	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/entity/common"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
	"github.com/go-redsync/redsync/v4"
)

type OrderService struct {
	repo           *pg.OrderRepo
	productService ProdService
	cartService    CartService
	redis          rds
	cancelFuncs    sync.Map
}

func NewOrderService(repo *pg.OrderRepo, productService ProdService, cartService CartService, redis rds) *OrderService {
	return &OrderService{
		repo:           repo,
		productService: productService,
		cartService:    cartService,
		redis:          redis,
		cancelFuncs:    sync.Map{},
	}
}

type ProdService interface {
	DecreaseProductCountFromOrder(ctx context.Context, prIds []uint, prsToOrder []entity.ProductsToOrder) error
	GetProductsByIds(productsId []uint) (products []entity.Product, err error)
}

type CartService interface {
	DeleteProductsFromCart(ctx context.Context, prIds []uint) error
}

type rds interface {
	LockProducts(productsIds []uint) ([]*redsync.Mutex, error)
	UnlockProducts(mutex []*redsync.Mutex) error
}

func (o *OrderService) TempOrder(ctx context.Context, newOrder []entity.ProductsToOrder) (orderId common.Id, err error) {
	newTempOrder := entity.Order{
		OrderId:     fmt.Sprintf("%v_%v", ctxpkg.UserFromCtx(ctx), time.Now().Format(time.DateTime)),
		Products:    newOrder,
		PaidExpired: false,
		Paid:        false,
		InsertedBy:  ctxpkg.UserFromCtx(ctx),
		Inserted:    time.Now(),
	}

	productIds := make([]uint, len(newOrder))
	for _, product := range newOrder {
		newTempOrder.OrderPrice += product.TotalPriceOnCount
		productIds = append(productIds, product.ProductID)
	}

	// TODO: анализ кейса когда lock не будет получен спустя ретраи

	mutex, err := o.redis.LockProducts(productIds)
	if err != nil {
		return common.Id{}, fmt.Errorf("create redis lock err: %w", err)
	}

	// проверка количества товаров
	if err := o.productsInStock(productIds, newTempOrder); err != nil {
		return common.Id{}, fmt.Errorf("inStock check error: %w", err)
	}

	// создать ордер
	tx, err := o.repo.CreateOrder(newTempOrder)
	if err != nil {
		return common.Id{}, fmt.Errorf("error when create temp order: %w", err)
	}

	// уменьшить кол-во товаров. В случае ошибки - роллбек ордера
	if err := o.productService.DecreaseProductCountFromOrder(ctx, productIds, newOrder); err != nil {
		if err := tx.Rollback().Error; err != nil {
			return common.Id{}, fmt.Errorf("create temp order error: %q, rollback decrease product error: %q", err, err)
		}
		log.Println("successful rollback DecreaseProductCountFromOrder")
		return common.Id{}, fmt.Errorf("error when decrease prods in stock: %w", err)
	}

	// удалить товары из корзины после создания заказа
	if err := o.cartService.DeleteProductsFromCart(ctx, productIds); err != nil {
		log.Printf("error delete products from cart: %v\n", err)
	}

	if err := o.redis.UnlockProducts(mutex); err != nil {
		log.Printf("redis unlock error: %v\n", err)
	}

	// заказ падает во временный, ожидаем оплаты
	ctxTimer, cancel := context.WithCancel(context.Background())
	o.cancelFuncs.Store(newTempOrder.OrderId, cancel)
	go o.paymentWait(ctxTimer, newTempOrder.OrderId)

	return common.Id{Id: newTempOrder.OrderId}, nil
}

func (o *OrderService) productsInStock(productIds []uint, newTempOrder entity.Order) error {
	products, err := o.productService.GetProductsByIds(productIds)
	if err != nil {
		return fmt.Errorf("get products error: %w", err)
	}

	for _, orderProduct := range newTempOrder.Products {
		for _, product := range products {
			if orderProduct.ProductID == product.ProductID {
				if orderProduct.CountInOrder > product.Count {
					return fmt.Errorf("not enough quantity for product %q. To order: %v, available: %v", orderProduct.Name, orderProduct.CountInOrder, product.Count)
				}
			}
		}
	}
	return nil
}

func (o *OrderService) MarkPaid(orderId string) error {
	if err := o.repo.MarkPaid(orderId); err != nil {
		return fmt.Errorf("mark paid orderId %v error: %w", orderId, err)
	}

	if cancel, ok := o.cancelFuncs.Load(orderId); ok {
		if ctxCancel, ok := cancel.(context.CancelFunc); ok {
			ctxCancel()
			o.cancelFuncs.Delete(orderId)
		}
	}
	return nil
}

func (o *OrderService) paymentWait(ctx context.Context, orderId string) {
	select {
	case <-time.After(15 * time.Minute):
		order, err := o.repo.GetOrderById(orderId)
		if err != nil {
			log.Printf("GetOrderById error: orderId %v, err: %v,\n", orderId, err)
		}

		if !order.Paid {
			if err := o.repo.MarkExpired(orderId); err != nil {
				log.Printf("MarkExpired error: orderId %v, err: %v\n", orderId, err)
			} else {
				log.Printf("order %v mark expired after 15 mins\n", orderId)
			}
		}
		o.cancelFuncs.Delete(orderId)
	case <-ctx.Done():
		log.Printf("orderId %v was paid\n", orderId)
	}
}

func (o *OrderService) GetOrderById(ctx context.Context, orderId string) (order entity.Order, err error) {
	return o.repo.GetOrderById(orderId)
}
