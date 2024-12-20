package purchase

import (
	"context"
	"exclusive-web/web/dto"
	"exclusive-web/web/entity"
)

type ControllerPurchase struct {
	usecase PurchaseUsecaseInteface
}

func (ctrl ControllerPurchase) Create(ctx context.Context, purchase entity.CreatePurchase, userId string) (dto.BaseResponse, error) {
	return ctrl.usecase.Create(ctx, purchase, userId)
}
