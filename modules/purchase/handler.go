package purchase

import (
	"exclusive-web/web/dto"
	"exclusive-web/web/entity"
	"exclusive-web/web/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PurchaseRequestHandler struct {
	db   *gorm.DB
	ctrl ControllerPurchase
}

func NewRequestHandler(db *gorm.DB) PurchaseRequestHandler {
	return PurchaseRequestHandler{
		db: db,
	}
}

func (handler PurchaseRequestHandler) HandlerPurchase(router *gin.Engine) {
	purchaserepo := repository.NewPurchase(handler.db)
	usecase := PurchaseUsecase{
		repository: purchaserepo,
	}
	handler.ctrl = ControllerPurchase{
		usecase: usecase,
	}

	purchaseGroup := router.Group("/purchase")
	purchaseGroup.POST("/", handler.Create)
}

func (handler PurchaseRequestHandler) Create(ctx *gin.Context) {
	var err error
	purchase := []entity.Purchase{}

	err = ctx.BindJSON(&purchase)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	response, err := handler.ctrl.Create(ctx, purchase)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
