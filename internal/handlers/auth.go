package handlers

import (
	"net/http"

	"github.com/Archiker-715/e-commerce-api/internal/auth"
	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/errs"
	"github.com/Archiker-715/e-commerce-api/pkg/httpsrv"
)

type AuthHandler struct {
	auth *auth.AuthService
}

func NewAuthHandler(service auth.AuthService) *AuthHandler {
	return &AuthHandler{auth: &service}
}

func (a *AuthHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	var authData entity.UserAuthRegistration
	if err := httpsrv.JsonDecode(w, r, &authData, 0); err != nil {
		errs.WriteError(w, 0, http.StatusBadRequest, "failed to parse json")
		return
	}

}
