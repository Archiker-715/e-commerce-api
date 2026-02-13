package pg

import (
	"fmt"

	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg/query"
	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{DB: db}
}

func (p *OrderRepo) CreateOrder(o entity.Order) error {
	if err := p.DB.Raw(query.CreateOrder(),
		o.OrderId,
		o.InsertedBy,
		o.OrderPrice,
		o.Products,
		o.Temp,
		o.InsertedBy,
		o.Inserted,
	).Error; err != nil {
		return fmt.Errorf("DB err: %w", err)
	}
	return nil
}
