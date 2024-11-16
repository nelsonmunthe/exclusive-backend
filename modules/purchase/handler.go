package purchase

import (
	"exclusive-web/web/dto"
	"exclusive-web/web/entity"
	"exclusive-web/web/middleware"
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
	purchaseGroup.POST("/create", middleware.Authenticate(), handler.Create)
}

func (handler PurchaseRequestHandler) Create(ctx *gin.Context) {
	var err error
	authData, err := middleware.GetAuthDataStruct(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	purchase := entity.CreatePurchase{}

	err = ctx.BindJSON(&purchase)

	for index, _ := range purchase.Purchase {
		purchase.Purchase[index].Customer_id = authData.UserID
		// purchase.Purchase[index].Product_id = authData.UserID
	}

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
