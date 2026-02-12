package pg

import (
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg/query"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRulesRepo struct {
	DB *gorm.DB
}

func NewProductRulesRepo(db *gorm.DB) *ProductRulesRepo {
	return &ProductRulesRepo{DB: db}
}

func (p *ProductRulesRepo) PermOnProduct(userId uuid.UUID, productId uint) (hasPerm bool, err error) {
	if err = p.DB.Raw(query.PermOnProduct(), userId, productId).Scan(&hasPerm).Error; err != nil {
		return false, err
	}
	return
}

func (p *ProductRulesRepo) UserInMarket(userId uuid.UUID, marketId uint) (inMarket bool, err error) {
	if err = p.DB.Raw(query.UserInMarket(), userId, marketId).Scan(&inMarket).Error; err != nil {
		return false, err
	}
	return
}

func (p *ProductRulesRepo) AdminRole(userId uuid.UUID) (adm bool, err error) {
	if err = p.DB.Raw(query.AdminRole(), userId).Scan(&adm).Error; err != nil {
		return false, err
	}
	return
}
