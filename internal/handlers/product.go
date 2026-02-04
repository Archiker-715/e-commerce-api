package handlers

import (
	"net/http"

	uc "github.com/Archiker-715/e-commerce-api/internal/usecase"
)

type ProductHandler struct {
	product *uc.ProductService
}

func NewProductHandler(service uc.ProductService) *ProductHandler {
	return &ProductHandler{product: &service}
}

func (p *ProductHandler) GetProduct(w http.ResponseWriter, r http.Request) {
	// ctx := r.Context()
	query := r.URL.Query()
	productId, article := query.Get("productId"), query.Get("article")

	if productId != "" {

	} else if article != "" {
	}
}
