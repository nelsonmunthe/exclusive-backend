package repository

import (
	"context"
	"exclusive-web/web/entity"

	"gorm.io/gorm"
)

type Category struct {
	db *gorm.DB
}

func NewCategory(db *gorm.DB) Category {
	return Category{
		db: db,
	}
}

type CategoryRepositoryUsecase interface {
	FindById(ctx context.Context, id int) (entity.Category, error)
	FindAll(ctx context.Context) ([]entity.Category, error)
}

func (category Category) FindById(ctx context.Context, id int) (entity.Category, error) {
	categoryDetail := &entity.Category{}
	err := category.db.First(&categoryDetail, id).Error
	return *categoryDetail, err
}

func (category Category) FindAll(ctx context.Context) ([]entity.Category, error) {
	categories := make([]entity.Category, 0)

	err := category.db.WithContext(ctx).Find(&categories).Error
	return categories, err
}
