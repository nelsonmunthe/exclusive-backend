package product

import (
	"exclusive-web/web/dto"
	"exclusive-web/web/entity"
	"exclusive-web/web/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRequestHandler struct {
	db   *gorm.DB
	ctrl ProductController
}

func NewRequestHandler(db *gorm.DB) ProductRequestHandler {
	return ProductRequestHandler{
		db: db,
	}
}

func (handler ProductRequestHandler) HandlerProduct(router *gin.Engine) {
	productRepository := repository.NewProduct(handler.db)
	productUsecase := ProductUsecase{
		repository: productRepository,
	}
	handler.ctrl = ProductController{
		usecase: productUsecase,
	}

	productGroup := router.Group("/product")
	productGroup.GET("/all", handler.GetAllProduct)
	productGroup.GET("/flash-sell", handler.FindProductFlashSell)
	productGroup.GET("/best-product", handler.GetBestProduct)
	productGroup.POST("/create", handler.Create)
	productGroup.POST("/upload", handler.Upload)
	productGroup.GET("/detail/:productId", handler.Detail)
}

func (handler ProductRequestHandler) FindProductFlashSell(ctx *gin.Context) {
	response, err := handler.ctrl.FindProductFlashSell(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (handler ProductRequestHandler) GetAllProduct(ctx *gin.Context) {
	var err error
	paginate := dto.PaginationRequest{}

	err = ctx.BindQuery(&paginate)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	response, err := handler.ctrl.GetAllProduct(ctx, paginate)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (handler ProductRequestHandler) GetBestProduct(ctx *gin.Context) {
	var err error

	bestProduct := entity.BestProduct{}

	err = ctx.BindQuery(&bestProduct)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	response, err := handler.ctrl.GetBestProduct(ctx, bestProduct)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (handler ProductRequestHandler) Create(ctx *gin.Context) {

	var newProduct = entity.Product{}

	isInvalid := ctx.BindJSON(&newProduct)
	newProduct.Created_at = time.Now()
	newProduct.Updated_at = time.Now()

	if isInvalid != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(isInvalid.Error()))
		return
	}

	response, err := handler.ctrl.Create(ctx, newProduct)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (handler ProductRequestHandler) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("image")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	newName := uuid.New().String()
	fileName := newName + "-" + file.Filename
	filePath := "./assets/images/" + fileName

	err = ctx.SaveUploadedFile(file, filePath)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	response := dto.BaseResponse{
		Data:         "/static/" + fileName,
		Success:      true,
		MessageTitle: "Succeeded",
		Message:      "Create new product Succeeded",
	}

	ctx.JSON(http.StatusOK, response)
}

func (handler ProductRequestHandler) Detail(ctx *gin.Context) {
	productId, err := strconv.ParseInt(ctx.Param("productId"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	response, err := handler.ctrl.Detail(ctx, int(productId))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
