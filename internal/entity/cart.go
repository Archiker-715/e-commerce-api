package entity

import "github.com/google/uuid"

type Cart struct {
	User  uuid.UUID `gorm:"column:user_id"`
	Name  string    `json:"name" gorm:"column:name"`
	Price uint64    `json:"price" gorm:"column:price"`
	Count uint64    `json:"count" gorm:"column:count"`
}
