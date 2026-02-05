package pg

import (
	"fmt"

	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg/query"
	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (p *ProductRepo) GetProducts() (products []entity.Product, err error) {
	if err = p.DB.Raw(query.GetProduct()).Scan(&products).Error; err != nil {
		return []entity.Product{}, fmt.Errorf("DB err: %w", err)
	}
	return
}

func (p *ProductRepo) GetProductById(productId uint) (products []entity.Product, err error) {
	if err = p.DB.Raw(query.GetProductById(), productId).Scan(&products).Error; err != nil {
		return []entity.Product{}, fmt.Errorf("DB err: %w", err)
	}
	return
}

func (p *ProductRepo) GetProductByArticle(article uint) (products []entity.Product, err error) {
	if err = p.DB.Raw(query.GetProductByArticle(), article).Scan(&products).Error; err != nil {
		return []entity.Product{}, fmt.Errorf("DB err: %w", err)
	}
	return
}
