package category

import (
	"exclusive-web/web/dto"
	"exclusive-web/web/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryRequestHandler struct {
	db   *gorm.DB
	ctrl CategoryController
}

func NewRequestHandler(db *gorm.DB) CategoryRequestHandler {
	return CategoryRequestHandler{
		db: db,
	}
}

func (handler CategoryRequestHandler) HandlerRepository(router *gin.Engine) {
	categoryRepo := repository.NewCategory(handler.db)
	UsecaseCategory := CategoryUsecase{
		repository: categoryRepo,
	}
	handler.ctrl = CategoryController{
		usecase: UsecaseCategory,
	}

	categoryGroup := router.Group("/category")
	categoryGroup.GET("/all", handler.FindAll)
	categoryGroup.GET("/detail/:id", handler.FindById)
}

func (handler CategoryRequestHandler) FindById(ctx *gin.Context) {
	categoryId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	response, err := handler.ctrl.FindById(ctx, int(categoryId))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (handler CategoryRequestHandler) FindAll(ctx *gin.Context) {
	response, err := handler.ctrl.FindAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
