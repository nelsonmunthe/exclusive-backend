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

func (pc ProductController) GetBestProduct(ctx context.Context, bestProduct entity.BestProduct) (dto.BaseResponse, error) {
	return pc.usecase.GetBestProduct(ctx, bestProduct)
}

func (pc ProductController) Create(ctx context.Context, product entity.Product) (dto.BaseResponse, error) {
	return pc.usecase.Create(ctx, product)
}

func (pc ProductController) Detail(ctx context.Context, productId int) (dto.BaseResponse, error) {
	return pc.usecase.Detail(ctx, productId)
}

func (pc ProductController) GetAllProduct(ctx context.Context, paginate dto.PaginationRequest) (dto.BaseResponse, error) {
	return pc.usecase.GetAllProduct(ctx, paginate)
}
