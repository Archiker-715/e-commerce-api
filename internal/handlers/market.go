package handlers

import (
	"fmt"
	"net/http"

	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/errs"
	uc "github.com/Archiker-715/e-commerce-api/internal/usecase"
	"github.com/Archiker-715/e-commerce-api/pkg/httpsrv"
)

type MarketHandler struct {
	market *uc.MarketService
}

func NewMarketHandler(service *uc.MarketService) *MarketHandler {
	return &MarketHandler{market: service}
}

func (m *MarketHandler) AddMarket(w http.ResponseWriter, r *http.Request) {
	var market entity.Market
	if err := httpsrv.JsonDecode(w, r, &market, 0); err != nil {
		errs.WriteError(w, 0, http.StatusBadRequest, fmt.Sprintf("failed to parse input: %v", err))
		return
	}

	ctx := r.Context()
	if err := m.market.AddMarket(ctx, market.MarketName); err != nil {
		errs.WriteError(w, 0, http.StatusInternalServerError, fmt.Sprintf("failed to create market: %v", err))
		return
	}
	fmt.Fprintln(w, "OK")
}

func (m *MarketHandler) LinkUserMarket(w http.ResponseWriter, r *http.Request) {
	var link entity.LinkUserMarket
	if err := httpsrv.JsonDecode(w, r, &link, 0); err != nil {
		errs.WriteError(w, 0, http.StatusBadRequest, fmt.Sprintf("failed to parse input: %v", err))
		return
	}

	ctx := r.Context()
	if err := m.market.LinkUserMarket(ctx, link); err != nil {
		errs.WriteError(w, 0, http.StatusInternalServerError, fmt.Sprintf("failed to create market: %v", err))
		return
	}
	fmt.Fprintln(w, "OK")
}
