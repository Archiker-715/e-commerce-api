package entity

import "github.com/google/uuid"

type UserCart struct {
	User       uuid.UUID `gorm:"column:user_id"`
	ProductId  uint      `gorm:"column:product_id"`
	Name       string    `json:"name" gorm:"column:name"`
	Count      uint64    `json:"count" gorm:"column:count"`
	UnitPrice  uint64    `json:"unitPrice" gorm:"column:unitPrice"`
	TotalPrice uint64    `json:"totalPrice" gorm:"column:totalPrice"`
}
