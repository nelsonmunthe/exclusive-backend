package repository

import (
	"context"
	"exclusive-web/web/entity"

	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func New(db *gorm.DB) User {
	return User{
		db: db,
	}
}

type UserInterfaceRepository interface {
	FindById(ctx context.Context, userId uint) (entity.User, error)
	FindByUsername(ctx context.Context, username string) (entity.User, error)
	Create(ctx context.Context, user entity.User) (entity.User, error)
}

func (usr User) FindById(ctx context.Context, userId uint) (entity.User, error) {
	userDetail := entity.User{}
	err := usr.db.First(&userDetail, userId).Error
	return userDetail, err
}

func (usr User) FindByUsername(ctx context.Context, username string) (entity.User, error) {
	userDetail := entity.User{}
	err := usr.db.Model(&userDetail).Where("username = ?", username).First(&userDetail).Error
	return userDetail, err
}

func (usr User) Create(ctx context.Context, user entity.User) (entity.User, error) {
	err := usr.db.WithContext(ctx).Save(&user).Error
	return user, err
}
