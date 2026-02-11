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

func (m *MarketRepo) AddMarket(userId uuid.UUID, marketName string) error {
	if err := m.DB.Raw(query.AddMarket(),
		marketName,
		userId,
		time.Now(),
		userId,
	).Error; err != nil {
		return err
	}
	return nil
}

func (m *MarketRepo) LinkUserMarket(userIdToLink uuid.UUID, marketId uint) error {
	if err := m.DB.Raw(query.LinkUserMarket(),
		userIdToLink,
		marketId,
		userIdToLink,
		time.Now(),
	).Error; err != nil {
		return err
	}
	return nil
}

func (m *MarketRepo) CheckOwner(userFromCtx uuid.UUID, marketId uint) (bool, error) {
	var owner bool
	if err := m.DB.Raw(query.CheckOwner(), userFromCtx, marketId).Scan(&owner).Error; err != nil {
		return false, err
	}
	return owner, nil
}
