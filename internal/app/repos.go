package app

import (
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
	"gorm.io/gorm"
)

type repositories struct {
	AuthRepo         *pg.AuthRepo
	ProductRepo      *pg.ProductRepo
	ProductRulesRepo *pg.ProductRulesRepo
	UserCartRepo     *pg.UserCartRepo
	OrderRepo        *pg.OrderRepo
	MarketRepo       *pg.MarketRepo
}

func newRepositories(db *gorm.DB) *repositories {
	return &repositories{
		AuthRepo:         pg.NewAuthRepo(db),
		ProductRepo:      pg.NewProductRepo(db),
		ProductRulesRepo: pg.NewProductRulesRepo(db),
		UserCartRepo:     pg.NewUserCartRepo(db),
		OrderRepo:        pg.NewOrderRepo(db),
		MarketRepo:       pg.NewMarketRepo(db),
	}
}
