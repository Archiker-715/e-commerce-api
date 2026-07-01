package uc

import "github.com/Archiker-715/e-commerce-api/internal/elastic"

type SearchService struct {
	es *elastic.Elastic
}

func NewSearchService(es *elastic.Elastic) *SearchService {
	return &SearchService{es: es}
}
