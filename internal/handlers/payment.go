package handlers

import (
	"net/http"

	uc "github.com/Archiker-715/e-commerce-api/internal/usecase"
)

type PaymentHandler struct {
	Payment *uc.PaymentService
}

func NewPaymentHandler(service *uc.PaymentService) *PaymentHandler {
	return &PaymentHandler{Payment: service}
}

func (o *PaymentHandler) DoPayment(w http.ResponseWriter, r *http.Request) {
	// TODO: доделать
}
