package auth

import (
	"context"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type tokenContextKey struct{}
type UserContextKey struct{}

var TokenCtxKey = tokenContextKey{}
var UserCtxKey = UserContextKey{}

func TokenToCtx(ctx context.Context, token *jwt.Token) context.Context {
	return context.WithValue(ctx, TokenCtxKey, token)
}

func TokenFromCtx(ctx context.Context) *jwt.Token {
	return ctx.Value(TokenCtxKey).(*jwt.Token)
}

func UserToCtx(ctx context.Context, userId uuid.UUID) context.Context {
	return context.WithValue(ctx, UserCtxKey, userId)
}

func UserFromCtx(ctx context.Context) uuid.UUID {
	return ctx.Value(UserCtxKey).(uuid.UUID)
}
