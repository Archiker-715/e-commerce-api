package uc

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Archiker-715/e-commerce-api/internal/auth"
	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/entity/common"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
	"gorm.io/gorm"
)

type ProductService struct {
	repo      *pg.ProductRepo
	rulesRepo pg.RulesRepo
}

func NewProductService(repo *pg.ProductRepo, rRepo pg.RulesRepo) *ProductService {
	return &ProductService{
		repo:      repo,
		rulesRepo: rRepo,
	}
}

type ProdService interface {
	DecreaseProductCountFromOrder(ctx context.Context, prIds []uint, prsToOrder []entity.ProductsToOrder) error
}

var forbiddenError error = errors.New("not enough rights")

func (p *ProductService) GetProduct(ctx context.Context, productId, article uint) ([]entity.Product, error) {
	ok, err := p.rulesRepo.PermOnProduct(auth.UserFromCtx(ctx), productId)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, forbiddenError
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

	var maxAttempts = 10
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

func (p *ProductService) UpdateProduct(ctx context.Context, productId uint, updPr entity.UpdateProduct) (common.Id, error) {
	var prId common.Id
	adm, err := p.rulesRepo.AdminRole(auth.UserFromCtx(ctx))
	if err != nil {
		return prId, err
	}
	if adm {
		if err := p.repo.UpdatePrice(productId, updPr.Price); err != nil {
			return prId, err
		}
		return common.Id{Id: productId}, nil
	}

	ok, err := p.rulesRepo.PermOnProduct(auth.UserFromCtx(ctx), productId)
	if err != nil {
		return prId, err
	}
	if !ok {
		return prId, forbiddenError
	}

	updateProduct := entity.Product{
		Name:        updPr.Name,
		Description: updPr.Description,
		Category:    updPr.Category,
		Price:       updPr.Price,
		Count:       updPr.Count,
		Active:      updPr.Active,
		Options:     updPr.Options,
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

func (p *ProductService) DecreaseProductCountFromOrder(ctx context.Context, prIds []uint, prsToOrder []entity.ProductsToOrder) error {
	products, err := p.repo.GetProductsByIds(prIds)
	if err != nil {
		return err
	}

	var sqlVals string
	for _, pr := range products {
		for _, prToOrd := range prsToOrder {
			if pr.Count < prToOrd.CountInOrder {
				return fmt.Errorf("not enough count product %q on stock. Available: %d, in order: %d", pr.Name, pr.Count, prToOrd.CountInOrder)
			}
			sqlVals += fmt.Sprintf("(%v, %v),", pr.ProductID, prToOrd.CountInOrder)
		}
	}
	sqlVals = strings.TrimRight(sqlVals, ",")

	return p.repo.DecreaseProductCountFromOrder(sqlVals)
}
