package product

import (
	"context"
	"exclusive-web/web/dto"
	"exclusive-web/web/entity"
	"exclusive-web/web/repository"
	"fmt"
)

type ProductUsecase struct {
	repository repository.ProductInterfaceRepository
}

type ProductInterfaceUsecase interface {
	FindProductFlashSell(ctx context.Context) (dto.BaseResponse, error)
	GetBestProduct(ctx context.Context) ([]dto.BaseResponse, error)
	Create(ctx context.Context, product entity.Product) (dto.BaseResponse, error)
}

func (pu ProductUsecase) FindProductFlashSell(ctx context.Context) (dto.BaseResponse, error) {
	products, err := pu.repository.FindProductFlashSell(ctx)
	if err != nil {
		return dto.DefaultErrorBaseResponseWithMessage(err)
	}

	return dto.BaseResponse{
		Data: products,
	}, err
}

func (pu ProductUsecase) GetBestProduct(ctx context.Context) (dto.BaseResponse, error) {
	products, err := pu.repository.GetBestProduct(ctx)
	if err != nil {
		return dto.DefaultErrorBaseResponseWithMessage(err)
	}

	return dto.BaseResponse{
		Data: products,
	}, err
}

func (pu ProductUsecase) Create(ctx context.Context, product entity.Product) (dto.BaseResponse, error) {
	fmt.Println("product", product)
	newProduct, err := pu.repository.Create(ctx, product)

	if err != nil {
		return dto.DefaultErrorBaseResponseWithMessage(err)
	}

	return dto.BaseResponse{
		Data:         newProduct,
		Success:      true,
		MessageTitle: "Succeeded",
		Message:      "Create new product Succeeded",
	}, nil
}
