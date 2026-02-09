package pg

import (
	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg/query"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartRepo struct {
	DB *gorm.DB
}

func NewCartRepo(db *gorm.DB) *CartRepo {
	return &CartRepo{DB: db}
}

func (c *CartRepo) GetUserCart(userId uuid.UUID) (userCart []entity.Cart, err error) {
	if err = c.DB.Raw(query.GetUserCart(), userId).Scan(&userCart).Error; err != nil {
		return []entity.Cart{}, err
	}
	return
}
