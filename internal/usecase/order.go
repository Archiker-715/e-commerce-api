package uc

import (
	"context"
	"fmt"
	"time"

	"github.com/Archiker-715/e-commerce-api/internal/auth"
	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
)

type OrderService struct {
	repo           *pg.OrderRepo
	productService ProdService
	cartService    CartService
}

func NewOrderService(repo *pg.OrderRepo, productService ProdService, cartService CartService) *OrderService {
	return &OrderService{
		repo:           repo,
		productService: productService,
		cartService:    cartService,
	}
}

func (o *OrderService) TempOrder(ctx context.Context, newOrder []entity.ProductsToOrder) error {
	newTempOrder := entity.Order{
		OrderId:    fmt.Sprintf("%v%v", auth.UserFromCtx(ctx), time.Now().Format(time.DateTime)),
		Products:   newOrder,
		Temp:       true,
		InsertedBy: auth.UserFromCtx(ctx),
		Inserted:   time.Now(),
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
		return fmt.Errorf("error when create temp order: %w", err)
	}

	if err := o.cartService.MarkOrdered(ctx, productIds, newTempOrder.OrderId); err != nil {
		return fmt.Errorf("error when mark products ordered: %w", err)
	}

	// TODO: делать роллбек если не удалось создать темп ордер

	// TODO: здесь нужен таймер на удаление заказа или отмена таймера при оплате

	return nil
}
