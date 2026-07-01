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
)

type OrderService struct {
	repo           *pg.OrderRepo
	productService ProdService
	cartService    CartService
	cancelFuncs    sync.Map
}

func NewOrderService(repo *pg.OrderRepo, productService ProdService, cartService CartService) *OrderService {
	return &OrderService{
		repo:           repo,
		productService: productService,
		cartService:    cartService,
		cancelFuncs:    sync.Map{},
	}
}

type ProdService interface {
	GetProductsByIds(productsId []uint) (products []entity.Product, err error)
	ReserveStock(ctx context.Context, newTempOrder entity.Order) error
	ConfirmReserve(ctx context.Context, productsToReserve []entity.ProductsInOrder) error
	DeclineReserve(ctx context.Context, productsToReserve []entity.ProductsInOrder) error
}

type CartService interface {
	DeleteProductsFromCart(ctx context.Context, prIds []uint) error
}

func (o *OrderService) TempOrder(ctx context.Context, newOrder []entity.ProductsInOrder) (orderId common.Id, err error) {
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

	// зарезервировать товары
	if err := o.productService.ReserveStock(ctx, newTempOrder); err != nil {
		return common.Id{}, fmt.Errorf("reserve stock error: %w", err)
	}

	// создать ордер
	if orderCreateErr := o.repo.CreateOrder(newTempOrder); orderCreateErr != nil {
		if declineResErr := o.productService.DeclineReserve(ctx, newTempOrder.Products); err != nil {
			return common.Id{}, fmt.Errorf("error when create temp order: %w, error when decline reserve: %w", orderCreateErr, declineResErr)
		}
		return common.Id{}, fmt.Errorf("error when create temp order: %w", err)
	}

	// удалить товары из корзины после создания заказа
	if err := o.cartService.DeleteProductsFromCart(ctx, productIds); err != nil {
		log.Printf("error delete products from cart: %v\n", err)
	}

	// заказ падает во временный, ожидаем оплаты
	ctxTimer, cancel := context.WithCancel(context.Background())
	o.cancelFuncs.Store(newTempOrder.OrderId, cancel)
	go o.paymentWait(ctxTimer, newTempOrder)

	return common.Id{Id: newTempOrder.OrderId}, nil
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

func (o *OrderService) paymentWait(ctx context.Context, newTempOrder entity.Order) {
	select {
	case <-time.After(15 * time.Minute):
		order, err := o.repo.GetOrderById(newTempOrder.OrderId)
		if err != nil {
			log.Printf("GetOrderById error: orderId %v, err: %v,\n", newTempOrder.OrderId, err)
		}

		if !order.Paid {
			if err := o.repo.MarkExpired(newTempOrder.OrderId); err != nil {
				log.Printf("MarkExpired error: orderId %v, err: %v\n", newTempOrder.OrderId, err)
			} else {
				log.Printf("order %v mark expired after 15 mins\n", newTempOrder.OrderId)
			}
		}
		o.cancelFuncs.Delete(newTempOrder.OrderId)
		if err := o.productService.DeclineReserve(ctx, newTempOrder.Products); err != nil {
			log.Printf("decline reserve error: %v\n", err)
		}
	case <-ctx.Done():
		log.Printf("orderId %v was paid\n", newTempOrder.OrderId)
		if err := o.productService.ConfirmReserve(ctx, newTempOrder.Products); err != nil {
			log.Printf("confirm reserve error: %v\n", err)
		}
	}
}

func (o *OrderService) GetOrderById(ctx context.Context, orderId string) (order entity.Order, err error) {
	return o.repo.GetOrderById(orderId)
}
