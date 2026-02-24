package uc

import (
	"context"
	"fmt"

	"github.com/Archiker-715/e-commerce-api/internal/kafka"
)

type PaymentService struct {
	orderService OrderSrv
	publisher    kafka.PaidOrderEventPublisher
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
