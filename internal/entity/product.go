package entity

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ProductID   uint                   `json:"productId"`
	MarketId    uint                   `json:"marketId"`
	Name        string                 `json:"name"`
	Description string                 `json:"descr" `
	Category    string                 `json:"category"`
	Price       uint                   `json:"price"`
	Count       uint                   `json:"count"`
	Active      bool                   `json:"active"`
	Options     map[string]interface{} `json:"options"`
	Article     string                 `json:"article"`
	InsertedBy  uuid.UUID              `json:"inserted_by"`
	Inserted    time.Time              `json:"inserted"`
	UpdatedBy   uuid.UUID              `json:"updated_by"`
	Updated     time.Time              `json:"updated"`
}

type CreateProduct struct {
	MarketId    uint                   `json:"marketId"`
	Name        string                 `json:"name"`
	Description string                 `json:"descr"`
	Category    string                 `json:"category"`
	Price       uint                   `json:"price"`
	Count       uint                   `json:"count"`
	Active      bool                   `json:"active"`
	Options     map[string]interface{} `json:"options,omitempty"`
}

type UpdateProduct struct {
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"descr,omitempty"`
	Category    string                 `json:"category,omitempty"`
	Price       uint                   `json:"price,omitempty"`
	Count       uint                   `json:"count,omitempty"`
	Active      bool                   `json:"active,omitempty"`
	Options     map[string]interface{} `json:"options,omitempty"`
}
