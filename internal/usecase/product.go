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
	PermOnProduct(userId uuid.UUID, productId uint) (hasPerm bool, err error)
	UserInMarket(userId uuid.UUID, marketId uint) (inMarket bool, err error)
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
	ok, err := p.rulesRepo.PermOnProduct(auth.UserFromCtx(ctx), productId)
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

	var productId common.Id
	ok, err := p.rulesRepo.UserInMarket(auth.UserFromCtx(ctx), pr.MarketId)
	if err != nil {
		return productId, err
	}
	if !ok {
		return productId, forbiddenError
	}

	newProduct := entity.Product{
		MarketId:    pr.MarketId,
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
		id, err := p.repo.CreateProduct(newProduct)
		if err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				continue
			}
			return productId, err
		}
		return common.Id{Id: id}, nil
	}
	return productId, errors.New("duplicate error when generate article")
}

func (p *ProductService) UpdateProduct(ctx context.Context, productId uint, pr entity.UpdateProduct) (common.Id, error) {
	var prId common.Id
	ok, err := p.rulesRepo.PermOnProduct(auth.UserFromCtx(ctx), productId)
	if err != nil {
		return prId, err
	}
	if !ok {
		return prId, forbiddenError
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
		return prId, err
	}
	return common.Id{Id: productId}, nil
}

func (p *ProductService) DeleteProduct(ctx context.Context, productId uint) error {
	ok, err := p.rulesRepo.PermOnProduct(auth.UserFromCtx(ctx), productId)
	if err != nil {
		return err
	}
	if !ok {
		return forbiddenError
	}
	return p.repo.DeleteProduct(productId)
}
