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

func (p *ProductRepo) GetProductsByIds(productId []uint) (products []entity.Product, err error) {
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

func (p *ProductRepo) CreateProduct(pr entity.Product) (productId uint, err error) {
	if err = p.DB.Raw(query.CreateProduct(),
		pr.Name,
		pr.Description,
		pr.Category,
		pr.Price,
		pr.Count,
		pr.Active,
		pr.Options,
		pr.Article,
		pr.InsertedBy,
		pr.Inserted,
		pr.MarketId,
	).Scan(&productId).Error; err != nil {
		return 0, fmt.Errorf("DB err: %w", err)
	}
	return
}

func (p *ProductRepo) UpdateProduct(pr entity.Product) error {
	if err := p.DB.Raw(query.UpdateProduct(),
		pr.Name,
		pr.Description,
		pr.Category,
		pr.Price,
		pr.Count,
		pr.Active,
		pr.Options,
		pr.UpdatedBy,
		pr.Updated,
		pr.ProductID,
	).Error; err != nil {
		return fmt.Errorf("DB err: %w", err)
	}
	return nil
}

func (p *ProductRepo) DeleteProduct(productId uint) error {
	if err := p.DB.Raw(query.DeleteProduct(), productId).Error; err != nil {
		return fmt.Errorf("DB err: %w", err)
	}
	return nil
}

func (p *ProductRepo) UpdatePrice(productId, price uint) error {
	if err := p.DB.Raw(query.UpdatePrice(), price, productId).Error; err != nil {
		return fmt.Errorf("DB err: %w", err)
	}
	return nil
}

func (p *ProductRepo) DecreaseProductCountFromOrder(decreaseProducts string) error {
	if err := p.DB.Raw(query.DecreaseProductCountFromOrder(), decreaseProducts).Error; err != nil {
		return fmt.Errorf("DB err: %w", err)
	}
	return nil
}

func (p *ProductRepo) IncreaseProductCountFromOrder(increaseProducts string) error {
	if err := p.DB.Raw(query.IncreaseProductCountFromOrder(), increaseProducts).Error; err != nil {
		return fmt.Errorf("DB err: %w", err)
	}
	return nil
}
