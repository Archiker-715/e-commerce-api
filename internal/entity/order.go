package entity

import (
	"time"

	"github.com/google/uuid"
)

type ProductsToOrder struct {
	ProductID         uint   `json:"productId"`
	Name              string `json:"name"`
	TotalPriceOnCount uint   `json:"totalPriceOnCount"`
	CountInOrder      uint   `json:"countInOrder"`
}

type Order struct {
	OrderId     string            `json:"orderId"`
	OrderPrice  uint              `json:"orderPrice"`
	Products    []ProductsToOrder `json:"products"`
	PaidExpired bool              `json:"paid_expired"`
	Paid        bool              `json:"paid"`
	InsertedBy  uuid.UUID         `json:"inserted_by"`
	Inserted    time.Time         `json:"inserted"`
	UpdatedBy   uuid.UUID         `json:"updated_by"`
	Updated     time.Time         `json:"updated"`
}
