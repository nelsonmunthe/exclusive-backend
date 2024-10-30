package user

import (
	"context"
	"exclusive-web/web/dto"
	"exclusive-web/web/entity"
)

type ControllerUser struct {
	usecase UserUsecaseInterface
}

func (ctr ControllerUser) FindById(ctx context.Context, userId uint) (dto.BaseResponse, error) {
	return ctr.usecase.FindById(ctx, userId)
}

func (ctr ControllerUser) Login(ctx context.Context, userLogin UserLogin) (dto.BaseResponse, error) {
	return ctr.usecase.Login(ctx, userLogin)
}

func (ctr ControllerUser) Create(ctx context.Context, user entity.User) (dto.BaseResponse, error) {
	return ctr.usecase.Create(ctx, user)
}
