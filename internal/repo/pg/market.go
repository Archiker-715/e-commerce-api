package pg

import (
	"time"

	"github.com/Archiker-715/e-commerce-api/internal/repo/pg/query"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MarketRepo struct {
	DB *gorm.DB
}

func NewMarketRepo(db *gorm.DB) *MarketRepo {
	return &MarketRepo{DB: db}
}

func (m *MarketRepo) AddMarket(userId, newMarketId uuid.UUID, marketName string) error {
	if err := m.DB.Raw(query.AddMarket(),
		newMarketId,
		marketName,
		userId,
		time.Now(),
		userId,
	).Error; err != nil {
		return err
	}
	return nil
}

func (m *MarketRepo) LinkUserMarket(userId uuid.UUID, marketId uuid.UUID) error {
	if err := m.DB.Raw(query.LinkUserMarket(),
		userId,
		marketId,
		userId,
		time.Now(),
	).Error; err != nil {
		return err
	}
	return nil
}
