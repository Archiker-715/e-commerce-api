package middleware

import (
	"net/http"
	"strings"

	"github.com/Archiker-715/e-commerce-api/internal/auth"
	ctxpkg "github.com/Archiker-715/e-commerce-api/internal/auth/ctx"
	"github.com/Archiker-715/e-commerce-api/internal/errs"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			errs.WriteError(w, 0, http.StatusUnauthorized, "empty token")
			return
		}

		token, err := auth.ParseToken(strings.TrimPrefix(authHeader, "Bearer "))
		if err != nil {
			errs.WriteError(w, 0, http.StatusUnauthorized, "invalid token")
			return
		}
		ctx := ctxpkg.UserToCtx(r.Context(), getUUID(token))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUUID(token *jwt.Token) uuid.UUID {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil
	}

	subClaim, ok := claims["user_id"].(string)
	if !ok {
		return uuid.Nil
	}

	userId, err := uuid.Parse(subClaim)
	if err != nil {
		return uuid.Nil
	}
	return userId
}
