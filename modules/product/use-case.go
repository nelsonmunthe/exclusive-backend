package product

import (
	"context"
	"exclusive-web/web/dto"
	"exclusive-web/web/entity"
	"exclusive-web/web/repository"
)

type ProductUsecase struct {
	repository repository.ProductInterfaceRepository
}

type ProductInterfaceUsecase interface {
	FindProductFlashSell(ctx context.Context) (dto.BaseResponse, error)
	GetBestProduct(ctx context.Context, bestProduct entity.BestProduct) ([]dto.BaseResponse, error)
	Create(ctx context.Context, product entity.Product) (dto.BaseResponse, error)
	Detail(ctx context.Context, productId int) (dto.BaseResponse, error)
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

func (pu ProductUsecase) GetBestProduct(ctx context.Context, bestProduct entity.BestProduct) (dto.BaseResponse, error) {
	products, err := pu.repository.GetBestProduct(ctx, bestProduct)
	if err != nil {
		return dto.DefaultErrorBaseResponseWithMessage(err)
	}

	return dto.BaseResponse{
		Data: products,
	}, err
}

func (pu ProductUsecase) Create(ctx context.Context, product entity.Product) (dto.BaseResponse, error) {
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

func (pu ProductUsecase) Detail(ctx context.Context, productId int) (dto.BaseResponse, error) {
	product, err := pu.repository.Detail(ctx, productId)
	if err != nil {
		return dto.DefaultErrorBaseResponseWithMessage(err)
	}

	return dto.BaseResponse{
		Data:         product,
		Success:      true,
		MessageTitle: "Succeeded",
		Message:      "Get  product Detail Succeeded",
	}, nil
}

func (pu ProductUsecase) GetAllProduct(ctx context.Context, pagination dto.PaginationRequest) (dto.BaseResponse, error) {
	products, count, err := pu.repository.GetAllProduct(ctx, pagination)

	if err != nil {
		return dto.DefaultErrorBaseResponseWithMessage(err)
	}

	response := dto.BaseResponsePagination{
		Data:  products,
		Total: count,
	}

	return dto.BaseResponse{
		Data:         response,
		Success:      true,
		MessageTitle: "Succeeded",
		Message:      "Get all Product Succeeded",
	}, nil
}
