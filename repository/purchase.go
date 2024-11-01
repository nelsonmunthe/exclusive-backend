package repository

import (
	"context"
	"exclusive-web/web/entity"

	"gorm.io/gorm"
)

type Purchase struct {
	db *gorm.DB
}

func NewPurchase(db *gorm.DB) Purchase {
	return Purchase{
		db: db,
	}
}

type PurcaseRepositoryInteface interface {
	Create(ctx context.Context, purchase []entity.Purchase) ([]entity.Purchase, error)
}

func (p Purchase) Create(ctx context.Context, purchase []entity.Purchase) ([]entity.Purchase, error) {
	err := p.db.Model(&entity.Purchase{}).Create(purchase).Error
	return purchase, err
}
