package repository

import (
	"context"
	"exclusive-web/web/entity"
	"exclusive-web/web/utils/pagination"
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) Product {
	return Product{
		db: db,
	}
}

type ProductInterfaceRepository interface {
	FindProductFlashSell(ctx context.Context) ([]entity.Product, error)
	GetBestProduct(ctx context.Context) ([]entity.Product, error)
	Create(ctx context.Context, product entity.Product) (entity.Product, error)
}

func (prdct Product) FindProductFlashSell(ctx context.Context) ([]entity.Product, error) {
	products := make([]entity.Product, 0)
	err := prdct.db.WithContext(ctx).Where("flash_sell= ?", true).Find(&products).Error
	return products, err
}

func (prdct Product) GetBestProduct(ctx context.Context) ([]entity.Product, error) {
	products := make([]entity.Product, 0)

	err := prdct.db.WithContext(ctx).Order("rating desc").Limit(4).Find(&products).Error
	return products, err
}

func (prdct Product) GetAllProduct(ctx context.Context, page int, per_page int) ([]entity.Product, error) {
	offset, limit, count := pagination.CountLimitAndOffset(page, per_page)
	fmt.Print(offset, limit, count)
	products := make([]entity.Product, 0)
	err := prdct.db.WithContext(ctx).Order("id desc").Offset(offset).Limit(limit).Find(&products).Error
	return products, err
}

func (prdct Product) Create(ctx context.Context, product entity.Product) (entity.Product, error) {
	err := prdct.db.WithContext(ctx).Save(&product).Error
	return product, err
}
