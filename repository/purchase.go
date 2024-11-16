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
	fmt.Println("purchase", purchase.Purchase[0])
	// err := p.db.Model(&entity.Purchase{}).Create(purchase.Purchase[0]).Error
	err := p.db.WithContext(ctx).Save(&purchase.Purchase).Error
	return purchase, err
}
