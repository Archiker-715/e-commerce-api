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

func (p *ProductRulesRepo) CheckPermission(userId uuid.UUID, permission string) (hasPerm bool, err error) {
	if err = p.DB.Raw(query.CheckPermission(), userId, permission).Scan(&hasPerm).Error; err != nil {
		return false, err
	}
	return
}
