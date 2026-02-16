package uc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Archiker-715/e-commerce-api/internal/auth"
	"github.com/Archiker-715/e-commerce-api/internal/entity"
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
	}
}

type OrderSrv interface {
	MarkPaid(ctx context.Context, orderId string) error
	GetOrderById(ctx context.Context, orderId string) (order entity.Order, err error)
}

func (o *OrderService) TempOrder(ctx context.Context, newOrder []entity.ProductsToOrder) error {
	newTempOrder := entity.Order{
		OrderId:     fmt.Sprintf("%v%v", auth.UserFromCtx(ctx), time.Now().Format(time.DateTime)),
		Products:    newOrder,
		PaidExpired: false,
		Paid:        false,
		InsertedBy:  auth.UserFromCtx(ctx),
		Inserted:    time.Now(),
	}

	productIds := make([]uint, len(newOrder))
	for _, product := range newOrder {
		newTempOrder.OrderPrice += product.TotalPriceOnCount
		productIds = append(productIds, product.ProductID)
	}

	if err := o.productService.DecreaseProductCountFromOrder(ctx, productIds, newOrder); err != nil {
		return fmt.Errorf("error when decrease prods in stock: %w", err)
	}

	if err := o.repo.CreateOrder(newTempOrder); err != nil {
		if rollbackErr := o.rollbackDecreaseProductCountFromOrder(ctx, newOrder); rollbackErr != nil {
			return fmt.Errorf("create temp order error: %q, rollback decrease product error: %q", err, rollbackErr)
		}
		return fmt.Errorf("error when create temp order: %w", err)
	}

	if err := o.cartService.DeleteProductsFromCart(ctx, productIds); err != nil {
		log.Printf("error delete products from cart: %v\n", err)
	}

	ctxTimer, cancel := context.WithCancel(context.Background())
	o.cancelFuncs.Store(newTempOrder.OrderId, cancel)

	go o.paymentWait(ctxTimer, newTempOrder.OrderId)

	return nil
}

func (o *OrderService) MarkPaid(ctx context.Context, orderId string) error {
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

func (o *OrderService) rollbackDecreaseProductCountFromOrder(ctx context.Context, newOrder []entity.ProductsToOrder) error {
	maxAttempts := 10
	for i := 1; i <= maxAttempts; i++ {
		if err := o.productService.IncreaseProductCountFromOrder(ctx, newOrder); err != nil {
			log.Println("rollbackDecreaseProductCountFromOrder error: ", err)
			if i == maxAttempts {
				return fmt.Errorf("failed to rollback product count after %d attempts: %w", maxAttempts, err)
			}
			continue
		}
		return nil
	}
	return errors.New("rollback attempts exhausted")
}

func (o *OrderService) GetOrderById(ctx context.Context, orderId string) (order entity.Order, err error) {
	return o.repo.GetOrderById(orderId)
}
