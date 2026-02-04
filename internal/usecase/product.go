package uc

import "github.com/Archiker-715/e-commerce-api/internal/repo/pg"

type ProductService struct {
	repo *pg.ProductRepo
}

func NewProductService(repo *pg.ProductRepo) *ProductService {
	return &ProductService{repo: repo}
}
