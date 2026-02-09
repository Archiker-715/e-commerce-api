package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId     uuid.UUID `gorm:"unique"`
	Login      string    `gorm:"unique"`
	Password   []byte
	InsertedBy uuid.UUID `json:"inserted_by" gorm:"column:inserted_by"`
	Inserted   time.Time `json:"inserted" gorm:"column:inserted"`
	UpdatedBy  uuid.UUID `json:"updated_by" gorm:"column:updated_by"`
	Updated    time.Time `json:"updated" gorm:"column:updated"`
}

type UserAuthRegistration struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
