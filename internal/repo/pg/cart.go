package pg

import (
	"github.com/Archiker-715/e-commerce-api/internal/entity"
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg/query"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserCartRepo struct {
	DB *gorm.DB
}

func NewUserCartRepo(db *gorm.DB) *UserCartRepo {
	return &UserCartRepo{DB: db}
}

func (c *UserCartRepo) GetUserCart(userId uuid.UUID) (userCart []entity.UserCart, err error) {
	if err = c.DB.Raw(query.GetUserCart(), userId).Scan(&userCart).Error; err != nil {
		return []entity.UserCart{}, err
	}
	return
}

func (c *UserCartRepo) AddProductToCart(productId uint, userId uuid.UUID) error {
	if err := c.DB.Raw(query.AddProductToCart(), userId, productId).Error; err != nil {
		return err
	}
	return nil
}

func (c *UserCartRepo) DeleteProductFromCart(productId uint, userId uuid.UUID) error {
	if err := c.DB.Raw(query.DeleteProductFromCart(), productId, userId).Error; err != nil {
		return err
	}
	return nil
}

func (c *UserCartRepo) IncreaseProductInCart(productId uint, userId uuid.UUID) error {
	if err := c.DB.Raw(query.IncreaseProductInCart(), productId, userId).Error; err != nil {
		return err
	}
	return nil
}

func (c *UserCartRepo) DecreaseProductInCart(productId uint, userId uuid.UUID) error {
	if err := c.DB.Raw(query.DecreaseProductInCart(), productId, userId, productId, userId).Error; err != nil {
		return err
	}
	return nil
}
