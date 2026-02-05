package uc

import (
	"context"

	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
)

type ProductService struct {
	repo *pg.ProductRepo
}

func NewProductService(repo *pg.ProductRepo) *ProductService {
	return &ProductService{repo: repo}
}

func (p *ProductService) GetProduct(ctx context.Context, productId, article uint) ([]entity.Product, error) {
	if productId != 0 {
		return p.repo.GetProductById(productId)
	} else if article != 0 {
		return p.repo.GetProductByArticle(article)
	}
	return p.repo.GetProducts()
}

func (p *ProductService) CreateProduct(ctx context.Context, product entity.CreateProduct) (productId int, err error) {

	// newProduct := entity.Product{
	// 	Name:        product.Name,
	// 	Description: product.Description,
	// 	Category:    product.Category,
	// 	Price:       product.Price,
	// 	Count:       product.Count,
	// 	Active:      product.Active,
	// 	Options:     product.Options,
	// 	// Article: ,
	// 	// InsertedBy: ,
	// 	// Inserted: ,
	// }
	return 0, nil
}
