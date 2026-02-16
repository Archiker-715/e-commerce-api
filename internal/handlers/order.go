package handlers

import (
	"fmt"
	"net/http"

	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/errs"
	uc "github.com/Archiker-715/e-commerce-api/internal/usecase"
	"github.com/Archiker-715/e-commerce-api/pkg/httpsrv"
)

type OrderHandler struct {
	order *uc.OrderService
}

func NewOrderHandler(service *uc.OrderService) *OrderHandler {
	return &OrderHandler{order: service}
}

func (o *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder []entity.ProductsToOrder
	if err := httpsrv.JsonDecode(w, r, &newOrder, 0); err != nil {
		errs.WriteError(w, 0, http.StatusBadRequest, fmt.Sprintf("failed to parse input: %v", err))
		return
	}

}
