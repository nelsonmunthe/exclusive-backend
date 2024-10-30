package category

import (
	"context"
	"exclusive-web/web/dto"
)

type CategoryController struct {
	usecase CategoryUsecase
}

func (ctrl CategoryController) FindById(ctx context.Context, id int) (dto.BaseResponse, error) {
	return ctrl.usecase.FindById(ctx, id)
}

func (ctrl CategoryController) FindAll(ctx context.Context) (dto.BaseResponse, error) {
	return ctrl.usecase.FindAll(ctx)
}
