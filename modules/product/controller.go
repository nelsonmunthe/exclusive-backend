package product

import (
	"context"
	"exclusive-web/web/dto"
	"exclusive-web/web/entity"
)

type ProductController struct {
	usecase ProductUsecase
}

func (pc ProductController) FindProductFlashSell(ctx context.Context) (dto.BaseResponse, error) {
	return pc.usecase.FindProductFlashSell(ctx)
}

func (pc ProductController) GetBestProduct(ctx context.Context) (dto.BaseResponse, error) {
	return pc.usecase.GetBestProduct(ctx)
}

func (pc ProductController) Create(ctx context.Context, product entity.Product) (dto.BaseResponse, error) {
	return pc.usecase.Create(ctx, product)
}
