package purchase

import (
	"context"
	"exclusive-web/web/dto"
	"exclusive-web/web/entity"
	"exclusive-web/web/repository"
)

type PurchaseUsecase struct {
	repository repository.PurcaseRepositoryInteface
}

type PurchaseUsecaseInteface interface {
	Create(ctx context.Context, purchase entity.CreatePurchase, userId string) (dto.BaseResponse, error)
}

func (uc PurchaseUsecase) Create(ctx context.Context, purchase entity.CreatePurchase, userId string) (dto.BaseResponse, error) {
	purchase, err := uc.repository.Create(ctx, purchase, userId)

	if err != nil {
		return dto.DefaultErrorBaseResponseWithMessage(err)
	}

	response := dto.BaseResponse{
		Data:         purchase,
		Success:      true,
		MessageTitle: "Succeeded",
		Message:      "Create Purchase Succeeded",
	}

	return response, err
}
