package uc

import (
	"context"
	"fmt"
)

type PaymentService struct {
	OrderService OrderSrv
}

func NewPaymentService(orderService OrderSrv) *PaymentService {
	return &PaymentService{
		OrderService: orderService,
	}
}

func (p *PaymentService) Payment(ctx context.Context, orderId string) error {
	// payment logic
	order, err := p.OrderService.GetOrderById(ctx, orderId)
	if err != nil {
		return fmt.Errorf("GetOrderById %v, error: %w", orderId, err)
	}
	// TODO: нужен гарант того, что если заказ оплачен, то он должен быть MarkPaid. Нужно подключать кафку
	if !order.PaidExpired {
		if err := p.OrderService.MarkPaid(ctx, orderId); err != nil {
			return fmt.Errorf("MarkPaid order %v, error: %w", orderId, err)
		}
	}
	return nil
}
