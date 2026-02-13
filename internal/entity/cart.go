package entity

import "github.com/google/uuid"

type UserCart struct {
	User       uuid.UUID
	ProductId  uint
	Name       string `json:"name"`
	Ordered    bool   `json:"ordered"`
	Count      uint64 `json:"count"`
	UnitPrice  uint64 `json:"unitPrice"`
	TotalPrice uint64 `json:"totalPrice"`
}
