package uc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	ctxpkg "github.com/Archiker-715/e-commerce-api/internal/auth/ctx"
	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/entity/common"
	"github.com/Archiker-715/e-commerce-api/internal/kafka"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
	kafkaGo "github.com/segmentio/kafka-go"
)

type OrderService struct {
	repo           *pg.OrderRepo
	productService ProdService
	cartService    CartService
	kafka          kafka.KafkaConsumer
	cancelFuncs    sync.Map
}

func NewOrderService(repo *pg.OrderRepo, productService ProdService, cartService CartService, kafka kafka.KafkaConsumer) *OrderService {
	return &OrderService{
		repo:           repo,
		productService: productService,
		cartService:    cartService,
		kafka:          kafka,
	}
}

type OrderSrv interface {
	MarkPaid(orderId string) error
	GetOrderById(ctx context.Context, orderId string) (order entity.Order, err error)
}

func (o *OrderService) TempOrder(ctx context.Context, newOrder []entity.ProductsToOrder) (orderId common.Id, err error) {
	newTempOrder := entity.Order{
		OrderId:     fmt.Sprintf("%v%v", ctxpkg.UserFromCtx(ctx), time.Now().Format(time.DateTime)),
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
	tx, err := o.productService.DecreaseProductCountFromOrder(ctx, productIds, newOrder)
	if err != nil {
		return common.Id{}, fmt.Errorf("error when decrease prods in stock: %w", err)
	}

	// создать ордер. В случае ошибки - роллбек
	if err := o.repo.CreateOrder(newTempOrder); err != nil {
		if err := tx.Rollback().Error; err != nil {
			return common.Id{}, fmt.Errorf("create temp order error: %q, rollback decrease product error: %q", err, err)
		}
		log.Println("successful rollback DecreaseProductCountFromOrder")
		return common.Id{}, fmt.Errorf("error when create temp order: %w", err)
	}

	// удалить товары из корзины после создания заказа
	if err := o.cartService.DeleteProductsFromCart(ctx, productIds); err != nil {
		log.Printf("error delete products from cart: %v\n", err)
	}

	// заказ падает во временный, ожидаем оплаты
	ctxTimer, cancel := context.WithCancel(context.Background())
	o.cancelFuncs.Store(newTempOrder.OrderId, cancel)
	go o.paymentWait(ctxTimer, newTempOrder.OrderId)

	return common.Id{Id: newTempOrder.OrderId}, nil
}

func (o *OrderService) ReadPaymentMessages() {
	handleMessage := func(m kafkaGo.Message) error {
		var paidOrder entity.Paid
		if err := json.Unmarshal(m.Value, &paidOrder); err != nil {
			log.Printf("unmashal err kafka message: key %v, partition %v. Err: %v\n", m.Key, m.Partition, err)
		}
		if err := o.MarkPaid(paidOrder.OrderId); err != nil {
			return err
		}
		return nil
	}

	if err := o.kafka.ReadMessage(handleMessage); err != nil {
		log.Printf("handle kafka message error: Err: %v\n", err)
	}
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
