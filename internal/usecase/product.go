package uc

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/Archiker-715/e-commerce-api/internal/auth"
	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/entity/common"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
	"gorm.io/gorm"
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

func (p *ProductService) CreateProduct(ctx context.Context, pr entity.CreateProduct) (common.Id, error) {
	genArticle := func() string {
		r := rand.Intn(25 + 1)
		upperR := byte('A' + r)
		lowerR := byte('a' + r)
		r2 := rand.Intn(10000000)
		return fmt.Sprintf("%v%v-%d", upperR, lowerR, r2)
	}

	newProduct := entity.Product{
		Name:        pr.Name,
		Description: pr.Description,
		Category:    pr.Category,
		Price:       pr.Price,
		Count:       pr.Count,
		Active:      pr.Active,
		Options:     pr.Options,
		InsertedBy:  auth.UserFromCtx(ctx),
	}

	var maxAttempts = 5
	for i := 0; i < maxAttempts; i++ {
		newProduct.Article = genArticle()
		newProduct.Inserted = time.Now()
		productId, err := p.repo.CreateProduct(newProduct)
		if err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				continue
			}
			return common.Id{}, err
		}
		return common.Id{Id: productId}, nil
	}
	return common.Id{}, errors.New("duplicate error when generate article")
}

func (p *ProductService) UpdateProduct(ctx context.Context, productId uint, pr entity.UpdateProduct) (common.Id, error) {
	updateProduct := entity.Product{
		Name:        pr.Name,
		Description: pr.Description,
		Category:    pr.Category,
		Price:       pr.Price,
		Count:       pr.Count,
		Active:      pr.Active,
		Options:     pr.Options,
		UpdatedBy:   auth.UserFromCtx(ctx),
		Updated:     time.Now(),
		ProductID:   productId,
	}

	if err := p.repo.UpdateProduct(updateProduct); err != nil {
		return common.Id{}, err
	}
	return common.Id{Id: productId}, nil
}
