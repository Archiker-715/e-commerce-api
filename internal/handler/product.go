package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/errs"
	uc "github.com/Archiker-715/e-commerce-api/internal/usecase"
	"github.com/Archiker-715/e-commerce-api/pkg/httpsrv"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	product *uc.ProductService
}

func NewProductHandler(service *uc.ProductService) *ProductHandler {
	return &ProductHandler{product: service}
}

var convertQueryParamError error = errors.New("convert to uint query parameters")
var emptyParamError error = errors.New("empty parameter")

func (p *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId, prodErr := toUint(r.URL.Query().Get("productId"))
	if errors.Is(prodErr, convertQueryParamError) {
		errs.WriteError(w, 0, http.StatusBadRequest, fmt.Sprintf("%v productId", prodErr))
		return
	}
	article, artErr := toUint(r.URL.Query().Get("article"))
	if errors.Is(artErr, convertQueryParamError) {
		errs.WriteError(w, 0, http.StatusInternalServerError, fmt.Sprintf("%v article", artErr))
		return
	}

	ctx := r.Context()
	products, err := p.product.GetProduct(ctx, uint(productId), uint(article))
	if err != nil {
		errs.WriteError(w, 0, http.StatusInternalServerError, fmt.Sprintf("get product: %v", err))
		return
	}

	httpsrv.JsonEncode(w, &products, 0)
}

func (p *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.CreateProduct
	if err := httpsrv.JsonDecode(w, r, &product, 0); err != nil {
		errs.WriteError(w, 0, http.StatusBadRequest, fmt.Sprintf("failed to parse input: %v", err))
		return
	}

	ctx := r.Context()
	productId, err := p.product.CreateProduct(ctx, product)
	if err != nil {
		errs.WriteError(w, 0, http.StatusInternalServerError, fmt.Sprintf("create product: %v", err))
		return
	}
	httpsrv.JsonEncode(w, &productId, 0)
}

func (p *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId, err := toUint(vars["productId"])
	if errors.Is(err, convertQueryParamError) || errors.Is(err, emptyParamError) {
		errs.WriteError(w, 0, http.StatusBadRequest, fmt.Sprintf("%v productId", err))
		return
	}

	var product entity.UpdateProduct
	if err := httpsrv.JsonDecode(w, r, &product, 0); err != nil {
		errs.WriteError(w, 0, http.StatusBadRequest, fmt.Sprintf("failed to parse input: %v", err))
		return
	}

	ctx := r.Context()
	updatedProd, err := p.product.UpdateProduct(ctx, productId, product)
	if err != nil {
		errs.WriteError(w, 0, http.StatusInternalServerError, fmt.Sprintf("Update product: %v", err))
		return
	}
	httpsrv.JsonEncode(w, &updatedProd, 0)
}

func (p *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId, err := toUint(vars["productId"])
	if errors.Is(err, convertQueryParamError) || errors.Is(err, emptyParamError) {
		errs.WriteError(w, 0, http.StatusBadRequest, fmt.Sprintf("%v productId", err))
		return
	}

	ctx := r.Context()
	if err := p.product.DeleteProduct(ctx, productId); err != nil {
		errs.WriteError(w, 0, http.StatusInternalServerError, fmt.Sprintf("Update product: %v", err))
		return
	}
	fmt.Fprintln(w, "OK")
}

func toUint(queryParam string) (uint, error) {
	if queryParam == "" {
		return 0, emptyParamError
	} else {
		u, err := strconv.Atoi(queryParam)
		if err != nil {
			return 0, fmt.Errorf("%w: %w", err, convertQueryParamError)
		}
		return uint(u), nil
	}
}
