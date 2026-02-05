package auth

import (
	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
)

type AuthService struct {
	repo *pg.AuthRepo
}

func NewAuthService(repo *pg.AuthRepo) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) Authorize(authData entity.UserAuthRegistration) {

}
