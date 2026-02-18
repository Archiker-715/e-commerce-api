package uc

import (
	"context"
	"errors"

	ctxpkg "github.com/Archiker-715/e-commerce-api/internal/auth/ctx"
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
	return c.repo.GetUserCart(ctxpkg.UserFromCtx(ctx))
}

func (c *UserCartService) GetProductsFromCartById(ctx context.Context, productsId []uint) ([]entity.UserCart, error) {
	return c.repo.GetProductsFromCartById(ctxpkg.UserFromCtx(ctx), productsId)
}

func (c *UserCartService) AddProductToCart(ctx context.Context, productId uint) error {
	return c.repo.AddProductToCart(productId, ctxpkg.UserFromCtx(ctx))
}

func (c *UserCartService) DeleteProductFromCart(ctx context.Context, productId uint) error {
	return c.repo.DeleteProductFromCart(productId, ctxpkg.UserFromCtx(ctx))
}

func (c *UserCartService) ChangeProductCount(ctx context.Context, productId uint, action string) error {
	switch action {
	case "increase":
		return c.repo.IncreaseProductInCart(productId, ctxpkg.UserFromCtx(ctx))
	case "decrease":
		return c.repo.DecreaseProductInCart(productId, ctxpkg.UserFromCtx(ctx))
	}
	return errors.New("change param is not in increase or decrease")
}

func (c *UserCartService) DeleteProductsFromCart(ctx context.Context, prIds []uint) error {
	return c.repo.DeleteProductsFromCart(ctxpkg.UserFromCtx(ctx), prIds)
}
