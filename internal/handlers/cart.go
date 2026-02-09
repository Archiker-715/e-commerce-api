package handlers

import (
	"fmt"
	"net/http"

	"github.com/Archiker-715/e-commerce-api/internal/errs"
	uc "github.com/Archiker-715/e-commerce-api/internal/usecase"
	"github.com/Archiker-715/e-commerce-api/pkg/httpsrv"
)

type CartHandler struct {
	cart *uc.CartService
}

func NewCartHandler(service *uc.CartService) *CartHandler {
	return &CartHandler{cart: service}
}

func (c *CartHandler) GetUserCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userCart, err := c.cart.GetUserCart(ctx)
	if err != nil {
		errs.WriteError(w, 0, http.StatusInternalServerError, fmt.Sprintf("%v productId", err))
		return
	}
	httpsrv.JsonEncode(w, &userCart, 0)
}

func (c *CartHandler) AddProductToCart(w http.ResponseWriter, r *http.Request) {

}
