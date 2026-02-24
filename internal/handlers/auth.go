package handlers

import (
	"fmt"
	"net/http"

	"github.com/Archiker-715/e-commerce-api/internal/auth"
	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/errs"
	"github.com/Archiker-715/e-commerce-api/pkg/httpsrv"
)

type AuthHandler struct {
	auth *auth.AuthService
}

func NewAuthHandler(service *auth.AuthService) *AuthHandler {
	return &AuthHandler{auth: service}
}

func (a *AuthHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	var authData entity.UserAuthRegistration
	if err := httpsrv.JsonDecode(w, r, &authData, 0); err != nil {
		errs.WriteError(w, 0, http.StatusBadRequest, "failed to parse json")
		return
	}

	authResp, err := a.auth.Authorize(authData)
	if err != nil {
		errs.WriteError(w, 0, http.StatusInternalServerError, fmt.Sprintf("auth error: %v", err))
		return
	}

	httpsrv.JsonEncode(w, &authResp, 0)
}

func (a *AuthHandler) Registration(w http.ResponseWriter, r *http.Request) {
	var registrationData entity.UserAuthRegistration
	if err := httpsrv.JsonDecode(w, r, &registrationData, 0); err != nil {
		errs.WriteError(w, 0, http.StatusBadRequest, "failed to parse json")
		return
	}

	err := a.auth.Registration(registrationData)
	if err != nil {
		errs.WriteError(w, 0, http.StatusInternalServerError, fmt.Sprintf("auth error: %v", err))
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Successfully sign Up! Now you can sign In via login & password")
}
