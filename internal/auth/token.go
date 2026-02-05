package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func (a *AuthService) generateToken(userId uuid.UUID) (entity.AuthResp, error) {
	expTokenTime := time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.MapClaims{
		"user_id": userId,
		"iat":     time.Now().Unix(),
		"exp":     expTokenTime,
		"jti":     uuid.New().String(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString(os.Getenv("JWT_KEY"))
	if err != nil {
		return entity.AuthResp{}, errors.New("signing token error")
	}

	out := entity.AuthResp{
		Token:     token,
		ExpiresIn: int(expTokenTime - time.Now().Unix()),
	}

	return out, nil
}

func ParseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %w", jwtToken.Header["alg"])
		}
		return os.Getenv("JWT_KEY"), nil
	})
}
