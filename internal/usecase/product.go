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
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RulesRepo interface {
	CheckPermission(userId uuid.UUID, permission string) (hasPerm bool, err error)
}

type ProductService struct {
	repo      *pg.ProductRepo
	rulesRepo RulesRepo
}

func NewProductService(repo *pg.ProductRepo, rRepo RulesRepo) *ProductService {
	return &ProductService{
		repo:      repo,
		rulesRepo: rRepo,
	}
}

var forbiddenError error = errors.New("not enough rights")

func (p *ProductService) GetProduct(ctx context.Context, productId, article uint) ([]entity.Product, error) {
	ok, err := p.rulesRepo.CheckPermission(auth.UserFromCtx(ctx), "READ")
	if err != nil {
		return []entity.Product{}, err
	}
	if !ok {
		return []entity.Product{}, forbiddenError
	}

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

	ok, err := p.rulesRepo.CheckPermission(auth.UserFromCtx(ctx), "INSERT")
	if err != nil {
		return common.Id{}, err
	}
	if !ok {
		return common.Id{}, forbiddenError
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
	ok, err := p.rulesRepo.CheckPermission(auth.UserFromCtx(ctx), "UPDATE")
	if err != nil {
		return common.Id{}, err
	}
	if !ok {
		return common.Id{}, forbiddenError
	}

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

func (p *ProductService) DeleteProduct(ctx context.Context, productId uint) error {
	ok, err := p.rulesRepo.CheckPermission(auth.UserFromCtx(ctx), "DELETE")
	if err != nil {
		return err
	}
	if !ok {
		return forbiddenError
	}
	return p.repo.DeleteProduct(productId)
}
