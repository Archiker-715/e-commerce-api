package uc

import (
	"context"
	"fmt"

	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/kafka"
)

type PaymentService struct {
	orderService OrderSrv
	publisher    kafka.PaidOrderEventPublisher
}

type OrderSrv interface {
	GetOrderById(ctx context.Context, orderId string) (order entity.Order, err error)
}

func NewPaymentService(orderService OrderSrv, publisher kafka.PaidOrderEventPublisher) *PaymentService {
	return &PaymentService{
		orderService: orderService,
		publisher:    publisher,
	}
}

func (p *PaymentService) Payment(ctx context.Context, orderId string) error {
	// payment logic
	order, err := p.orderService.GetOrderById(ctx, orderId)
	if err != nil {
		return fmt.Errorf("GetOrderById %v, error: %w", orderId, err)
	}

	p.publisher.PaidEventPublish(ctx, order.OrderId)

	return nil
}
