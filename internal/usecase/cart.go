package uc

import (
	"context"

	"github.com/Archiker-715/e-commerce-api/internal/auth"
	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
)

type CartService struct {
	repo *pg.CartRepo
}

func NewCartService(repo *pg.CartRepo) *CartService {
	return &CartService{repo: repo}
}

func (c *CartService) GetUserCart(ctx context.Context) ([]entity.Cart, error) {
	return c.repo.GetUserCart(auth.UserFromCtx(ctx))
}
