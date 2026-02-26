package handler

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
	var newOrder []entity.ProductsInOrder
	if err := httpsrv.JsonDecode(w, r, &newOrder, 0); err != nil {
		errs.WriteError(w, 0, http.StatusBadRequest, fmt.Sprintf("failed to parse input: %v", err))
		return
	}

	ctx := r.Context()
	orderId, err := o.order.TempOrder(ctx, newOrder)
	if err != nil {
		errs.WriteError(w, 0, http.StatusInternalServerError, fmt.Sprintf("failed to create temp order: %v", err))
		return
	}

	httpsrv.JsonEncode(w, &orderId, 0)
}
