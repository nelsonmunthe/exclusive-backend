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

	// Get a Tx for making transaction requests.
	tx := p.db.Begin()

	purchase_header := entity.Purchase_header{
		Total_price: float64(purchase.Total_price),
		Customer_id: userId,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}

	err := tx.Model(&purchase_header).Save(&purchase_header).Error

	if err != nil {
		tx.Rollback()
		return purchase, err
	}

	for index, _ := range purchase.Purchase_line {
		purchase.Purchase_line[index].Purchase_header_id = purchase_header.ID
		purchase.Purchase_line[index].Created_at = time.Now()
		purchase.Purchase_line[index].Updated_at = time.Now()
		isUpdateStockFailed := p.UpdateStockProduct(ctx, purchase.Purchase_line[index], tx)
		if isUpdateStockFailed != nil {
			tx.Rollback()
			return purchase, isUpdateStockFailed
		}
	}

	err = tx.WithContext(ctx).Save(&purchase.Purchase_line).Error

	if err != nil {
		tx.Rollback()
		return purchase, err
	}

	tx.Commit()
	return purchase, err
}

func (p Purchase) UpdateStockProduct(ctx context.Context, purchase_line entity.Purchase_line, db *gorm.DB) error {
	var err error
	product := entity.Product{}

	err = db.Model(&product).Where("id = ?", purchase_line.Product_id).First(&product).Error

	if err != nil {
		db.Rollback()
		return err
	}

	err = db.Model(&product).WithContext(ctx).
		Where("id = ?", product.ID).
		Update("total", product.Total-purchase_line.Quantity).
		Error

	return err
}
