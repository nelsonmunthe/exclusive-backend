package repository

import (
	"context"
	"exclusive-web/web/entity"
	"fmt"

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
	Create(ctx context.Context, purchase entity.CreatePurchase) (entity.CreatePurchase, error)
}

func (p Purchase) Create(ctx context.Context, purchase entity.CreatePurchase) (entity.CreatePurchase, error) {
	err := p.db.Model(&entity.Purchase{}).Create(purchase.Purchase).Error
	fmt.Println("purchase", purchase.Purchase)
	return purchase, err
}
