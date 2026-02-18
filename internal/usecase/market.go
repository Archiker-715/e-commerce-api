package uc

import (
	"context"
	"errors"

	ctxpkg "github.com/Archiker-715/e-commerce-api/internal/auth/ctx"
	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
)

type MarketService struct {
	repo *pg.MarketRepo
}

func NewMarketService(repo *pg.MarketRepo) *MarketService {
	return &MarketService{repo: repo}
}

func (m *MarketService) AddMarket(ctx context.Context, marketName string) error {
	return m.repo.AddMarket(ctxpkg.UserFromCtx(ctx), marketName)
}

func (m *MarketService) LinkUserMarket(ctx context.Context, link entity.LinkUserMarket) error {
	owner, err := m.repo.CheckOwner(ctxpkg.UserFromCtx(ctx), link.MarketId)
	if err != nil {
		return err
	}
	if !owner {
		return errors.New("forbidden request. Current user is not the market owner")
	}
	return m.repo.LinkUserMarket(link.UserToLink, link.MarketId)
}
