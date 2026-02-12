package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId     uuid.UUID
	Login      string
	Password   []byte
	InsertedBy uuid.UUID `json:"inserted_by"`
	Inserted   time.Time `json:"inserted"`
	UpdatedBy  uuid.UUID `json:"updated_by"`
	Updated    time.Time `json:"updated"`
}

type UserAuthRegistration struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
