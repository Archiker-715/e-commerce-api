package uc

import (
	"context"
	"errors"

	"github.com/Archiker-715/e-commerce-api/internal/auth"
	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
)

type UserCartService struct {
	repo *pg.UserCartRepo
}

func NewUserCartService(repo *pg.UserCartRepo) *UserCartService {
	return &UserCartService{repo: repo}
}

type CartService interface {
	DeleteProductsFromCart(ctx context.Context, prIds []uint) error
}

func (c *UserCartService) GetUserCart(ctx context.Context) ([]entity.UserCart, error) {
	return c.repo.GetUserCart(auth.UserFromCtx(ctx))
}

func (c *UserCartService) GetProductsFromCartById(ctx context.Context, productsId []uint) ([]entity.UserCart, error) {
	return c.repo.GetProductsFromCartById(auth.UserFromCtx(ctx), productsId)
}

func (c *UserCartService) AddProductToCart(ctx context.Context, productId uint) error {
	return c.repo.AddProductToCart(productId, auth.UserFromCtx(ctx))
}

func (c *UserCartService) DeleteProductFromCart(ctx context.Context, productId uint) error {
	return c.repo.DeleteProductFromCart(productId, auth.UserFromCtx(ctx))
}

func (c *UserCartService) ChangeProductCount(ctx context.Context, productId uint, action string) error {
	switch action {
	case "increase":
		return c.repo.IncreaseProductInCart(productId, auth.UserFromCtx(ctx))
	case "decrease":
		return c.repo.DecreaseProductInCart(productId, auth.UserFromCtx(ctx))
	}
	return errors.New("change param is not in increase or decrease")
}

func (c *UserCartService) DeleteProductsFromCart(ctx context.Context, prIds []uint) error {
	return c.repo.DeleteProductsFromCart(auth.UserFromCtx(ctx), prIds)
}
