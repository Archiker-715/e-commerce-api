package auth

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
	"gorm.io/gorm"
)

type AuthService struct {
	repo *pg.AuthRepo
}

func NewAuthService(repo *pg.AuthRepo) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) Authorize(authData entity.UserAuthRegistration) (entity.AuthResp, error) {
	hashed := sha256.Sum256([]byte(authData.Password))
	userId, err := a.repo.GetUserByLogPass(authData.Login, hashed[:])
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.AuthResp{}, errors.New("User does not exist or password is incorrect. Try again")
		}
		return entity.AuthResp{}, fmt.Errorf("get user error: %w", err)
	}
	return a.generateToken(userId)
}
