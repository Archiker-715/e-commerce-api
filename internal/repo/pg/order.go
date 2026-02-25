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

func (p *OrderRepo) CreateOrder(o entity.Order) (tx *gorm.DB, err error) {
	tx = p.DB.Raw(query.CreateOrder(),
		o.OrderId,
		o.InsertedBy,
		o.OrderPrice,
		o.Products,
		o.PaidExpired,
		o.Paid,
		o.InsertedBy,
		o.Inserted,
	)
	if tx.Error != nil {
		return nil, fmt.Errorf("DB err: %w", err)
	}
	return tx, nil
}

func (p *OrderRepo) GetOrderById(orderId string) (order entity.Order, err error) {
	if err = p.DB.Raw(query.GetOrderById(), orderId).Scan(&order).Error; err != nil {
		return entity.Order{}, fmt.Errorf("DB err: %w", err)
	}
	return
}

func (p *OrderRepo) MarkExpired(orderId string) error {
	if err := p.DB.Raw(query.MarkExpired(), orderId).Error; err != nil {
		return fmt.Errorf("DB err: %w", err)
	}
	return nil
}

func (p *OrderRepo) MarkPaid(orderId string) error {
	if err := p.DB.Raw(query.MarkPaid(), orderId).Error; err != nil {
		return fmt.Errorf("DB err: %w", err)
	}
	return nil
}
