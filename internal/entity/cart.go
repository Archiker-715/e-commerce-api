package entity

import "github.com/google/uuid"

type UserCart struct {
	User               uuid.UUID
	ProductId          uint
	Name               string `json:"name"`
	PurchaseAvailable  bool   `json:"purchaseAvailable"`
	NotAvailableReason error  `json:"notAvailiableReason"`
	Count              uint   `json:"count"`
	UnitPrice          uint   `json:"unitPrice"`
	TotalPrice         uint   `json:"totalPrice"`
}
