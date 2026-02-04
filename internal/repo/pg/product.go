package pg

import "gorm.io/gorm"

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}
