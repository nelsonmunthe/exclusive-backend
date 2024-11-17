package repository

import (
	"context"
	"exclusive-web/web/entity"
	"time"

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
	Create(ctx context.Context, purchase entity.CreatePurchase, userId string) (entity.CreatePurchase, error)
}

func (p Purchase) Create(ctx context.Context, purchase entity.CreatePurchase, userId string) (entity.CreatePurchase, error) {
	// create purchase header
	purchase_header := entity.Purchase_header{
		Total_price: float64(purchase.Total_price),
		Customer_id: userId,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}

	err := p.db.Model(&purchase_header).Save(&purchase_header).Error

	if err != nil {
		return purchase, err
	}

	for index, _ := range purchase.Purchase_line {
		purchase.Purchase_line[index].Purchase_header_id = purchase_header.ID
		purchase.Purchase_line[index].Created_at = time.Now()
		purchase.Purchase_line[index].Updated_at = time.Now()
	}

	err = p.db.WithContext(ctx).Save(&purchase.Purchase_line).Error
	return purchase, err
}
