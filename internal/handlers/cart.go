package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Archiker-715/e-commerce-api/internal/errs"
	uc "github.com/Archiker-715/e-commerce-api/internal/usecase"
	"github.com/Archiker-715/e-commerce-api/pkg/httpsrv"
	"github.com/gorilla/mux"
)

type CartHandler struct {
	cart *uc.UserCartService
}

func NewCartHandler(service *uc.UserCartService) *CartHandler {
	return &CartHandler{cart: service}
}

func (c *CartHandler) GetUserCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userCart, err := c.cart.GetUserCart(ctx)
	if err != nil {
		errs.WriteError(w, 0, http.StatusInternalServerError, fmt.Sprintf("%v", err))
		return
	}
	httpsrv.JsonEncode(w, &userCart, 0)
}

func (c *CartHandler) ChangeProductCount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]
	if action == "" && action != "increase" && action != "decrease" {
		errs.WriteError(w, 0, http.StatusBadRequest, "not enought args. Action must by increase or decrease")
		return
	}

	productId, err := toUint(vars["productId"])
	if errors.Is(err, convertQueryParamError) || errors.Is(err, emptyParamError) {
		errs.WriteError(w, 0, http.StatusBadRequest, fmt.Sprintf("%v productId", err))
		return
	}

	ctx := r.Context()
	if err := c.cart.ChangeProductCount(ctx, productId, action); err != nil {
		errs.WriteError(w, 0, http.StatusInternalServerError, fmt.Sprintf("%v", err))
		return
	}
	fmt.Fprintln(w, "OK")
}
