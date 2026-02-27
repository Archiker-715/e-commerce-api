package handler

import (
	"net/http"

	uc "github.com/Archiker-715/e-commerce-api/internal/usecase"
)

type SearchHandler struct {
	es *uc.SearchService
}

func NewSearchHandler(search *uc.SearchService) *SearchHandler {
	return &SearchHandler{es: search}
}

func (s *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {

}
