package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Archiker-715/e-commerce-api/internal/errs"
	uc "github.com/Archiker-715/e-commerce-api/internal/usecase"
	"github.com/Archiker-715/e-commerce-api/pkg/httpsrv"
)

type ProductHandler struct {
	product *uc.ProductService
}

func NewProductHandler(service uc.ProductService) *ProductHandler {
	return &ProductHandler{product: &service}
}

var convertQueryParamError error = errors.New("failed convert to uint query param")

func (p *ProductHandler) GetProduct(w http.ResponseWriter, r http.Request) {
	ctx := r.Context()
	productId, prodErr := toUint(r.URL.Query().Get("productId"))
	if errors.Is(prodErr, convertQueryParamError) {
		errs.WriteError(w, 0, http.StatusBadRequest, fmt.Sprintf("%v productId", prodErr))
		return
	}
	article, artErr := toUint(r.URL.Query().Get("article"))
	if errors.Is(artErr, convertQueryParamError) {
		errs.WriteError(w, 0, http.StatusBadRequest, fmt.Sprintf("%v article", artErr))
		return
	}

	products, err := p.product.GetProduct(ctx, uint(productId), uint(article))
	if err != nil {
		errs.WriteError(w, 0, http.StatusBadRequest, fmt.Sprintf("get product: %v", err))
		return
	}

	httpsrv.JsonEncode(w, &products, 0)
}

func toUint(input string) (uint, error) {
	if input == "" {
		return 0, nil
	} else {
		u, err := strconv.Atoi(input)
		if err != nil {
			return 0, convertQueryParamError
		}
		return uint(u), nil
	}
}
