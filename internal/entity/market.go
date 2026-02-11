package entity

import "github.com/google/uuid"

type Market struct {
	MarketName string `json:"marketName"`
}

type LinkUserMarket struct {
	MarketId   uint      `json:"marketId"`
	UserToLink uuid.UUID `json:"userToLink"`
}
