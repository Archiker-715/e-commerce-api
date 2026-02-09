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

func (c *UserCartService) GetUserCart(ctx context.Context) ([]entity.UserCart, error) {
	return c.repo.GetUserCart(auth.UserFromCtx(ctx))
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
