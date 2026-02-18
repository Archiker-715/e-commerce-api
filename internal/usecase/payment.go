package uc

import (
	"context"
	"encoding/json"
	"fmt"

	ctxpkg "github.com/Archiker-715/e-commerce-api/internal/auth/ctx"
	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/kafka"
)

type PaymentService struct {
	orderService OrderSrv
	kafka        kafka.KafkaProducer
}

func NewPaymentService(orderService OrderSrv, kafka kafka.KafkaProducer) *PaymentService {
	return &PaymentService{
		orderService: orderService,
		kafka:        kafka,
	}
}

func (p *PaymentService) Payment(ctx context.Context, orderId string) error {
	// payment logic
	order, err := p.orderService.GetOrderById(ctx, orderId)
	if err != nil {
		return fmt.Errorf("GetOrderById %v, error: %w", orderId, err)
	}

	jsonB, err := json.Marshal(entity.Paid{OrderId: orderId})
	if err != nil {
		return fmt.Errorf("marshal orderId err: %w", err)
	}
	if !order.PaidExpired {
		if err := p.kafka.SendMessage(ctxpkg.UserFromCtxAsStr(ctx), jsonB); err != nil {
			return err
		}
	}
	return nil
}
