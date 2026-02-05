package pg

import (
	"fmt"

	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg/query"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepo struct {
	DB *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{DB: db}
}

func (a *AuthRepo) GetUserByLogPass(login string, password []byte) (uuid.UUID, error) {
	var user entity.User
	if err := a.DB.Raw(query.GetUserByLogPass(), login, password).Scan(&user).Error; err != nil {
		return uuid.Nil, fmt.Errorf("DB err: %w", err)
	}
	return user.UserId, nil
}

func (a *AuthRepo) CreateUser(user entity.User) error {
	if err := a.DB.Raw(query.CreateUser(), user.UserId, user.Login, user.Password).Error; err != nil {
		return fmt.Errorf("DB err: %w", err)
	}
	return nil
}
