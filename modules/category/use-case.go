package category

import (
	"context"
	"exclusive-web/web/dto"
	"exclusive-web/web/repository"
)

type CategoryUsecase struct {
	repository repository.CategoryRepositoryUsecase
}

type CategoryIntefaceUsecase interface {
	FindById(ctx context.Context, id int) (dto.BaseResponse, error)
	FindAll(ctx context.Context) (dto.BaseResponse, error)
}

func (uc CategoryUsecase) FindById(ctx context.Context, id int) (dto.BaseResponse, error) {
	category, err := uc.repository.FindById(ctx, id)
	if err != nil {
		return dto.DefaultErrorBaseResponseWithMessage(err)
	}
	return dto.BaseResponse{
		Data: category,
	}, err
}

func (uc CategoryUsecase) FindAll(ctx context.Context) (dto.BaseResponse, error) {
	categories, err := uc.repository.FindAll(ctx)
	if err != nil {
		return dto.DefaultErrorBaseResponseWithMessage(err)
	}
	return dto.BaseResponse{
		Data: categories,
	}, err
}
