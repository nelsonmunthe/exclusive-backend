package repository

import (
	"context"
	"exclusive-web/web/dto"
	"exclusive-web/web/entity"
	"exclusive-web/web/utils/pagination"

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
	GetBestProduct(ctx context.Context, bestProduct entity.BestProduct) ([]entity.Product, error)
	Create(ctx context.Context, product entity.Product) (entity.Product, error)
	Detail(ctx context.Context, productId int) (entity.Product, error)
	GetAllProduct(ctx context.Context, paginate dto.PaginationRequest) ([]entity.Product, int64, error)
}

func (prdct Product) FindProductFlashSell(ctx context.Context) ([]entity.Product, error) {
	products := make([]entity.Product, 0)
	err := prdct.db.WithContext(ctx).Where("flash_sell= ?", true).Order("Created_at desc").Find(&products).Error
	return products, err
}

func (prdct Product) GetBestProduct(ctx context.Context, bestProduct entity.BestProduct) ([]entity.Product, error) {
	products := make([]entity.Product, 0)
	filter := "rating desc"
	if bestProduct.BestMonth {
		filter = "Created_at desc"
	}
	err := prdct.db.WithContext(ctx).Order(filter).Limit(4).Find(&products).Error
	return products, err
}

func (prdct Product) GetAllProduct(ctx context.Context, paginate dto.PaginationRequest) ([]entity.Product, int64, error) {
	var err error
	count := int64(0)
	offset, limit := pagination.CountLimitAndOffset(paginate.Page, paginate.PerPage)

	products := make([]entity.Product, 0)
	err = prdct.db.WithContext(ctx).Order("id desc").Offset(offset).Limit(limit).Find(&products).Error
	if err != nil {
		return products, count, err
	}

	err = prdct.db.WithContext(ctx).Table("products").Count(&count).Error
	return products, count, err
}

func (prdct Product) Create(ctx context.Context, product entity.Product) (entity.Product, error) {
	err := prdct.db.WithContext(ctx).Save(&product).Error
	return product, err
}

func (prdct Product) Detail(ctx context.Context, productId int) (entity.Product, error) {
	product := entity.Product{}
	err := prdct.db.First(&product, productId).Error

	return product, err
}
