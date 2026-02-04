package entity

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ProductID   uint                   `json:"id" gorm:"primaryKey;column:product_id"`
	Name        string                 `json:"name" gorm:"column:name"`
	Description string                 `json:"descr" gorm:"column:description"`
	Category    string                 `json:"category" gorm:"column:category"`
	Price       uint64                 `json:"price" gorm:"column:price"`
	Count       uint64                 `json:"count" gorm:"column:count"`
	Active      bool                   `json:"active" gorm:"column:active"`
	Options     map[string]interface{} `json:"options" gorm:"column:options;type:json"` // по идее сразу будет разобрано по json полям
	Article     int64                  `json:"article" gorm:"column:article;unique"`    // подумать над производительностью в части определения уникальности
	InsertedBy  uuid.UUID              `json:"inserted_by" gorm:"column:inserted_by"`
	Inserted    time.Time              `json:"inserted" gorm:"column:inserted"`
	UpdatedBy   uuid.UUID              `json:"updated_by" gorm:"column:updated_by"`
	Updated     time.Time              `json:"updated" gorm:"column:updated"`
}

type AddProduct struct {
	Name        string                 `json:"name"`
	Description string                 `json:"descr"`
	Category    string                 `json:"category"`
	Price       uint64                 `json:"price"`
	Count       uint64                 `json:"count"`
	Active      bool                   `json:"active"`
	Options     map[string]interface{} `json:"options,omitempty"`
}

type UpdateProduct struct {
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"descr,omitempty"`
	Category    string                 `json:"category,omitempty"`
	Price       uint64                 `json:"price,omitempty"`
	Count       uint64                 `json:"count,omitempty"`
	Active      bool                   `json:"active,omitempty"`
	Options     map[string]interface{} `json:"options,omitempty"`
}
