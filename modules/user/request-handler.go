package user

import (
	"exclusive-web/web/dto"
	"exclusive-web/web/entity"
	"exclusive-web/web/middleware"
	"exclusive-web/web/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRequestHandler struct {
	db  *gorm.DB
	ctr ControllerUser
}

func NewRequestHandler(db *gorm.DB) UserRequestHandler {
	return UserRequestHandler{
		db: db,
	}
}

func (handler UserRequestHandler) HandlerUser(router *gin.Engine) {
	userRepo := repository.New(handler.db)
	userUsecase := UserUsecase{
		repository: userRepo,
	}
	handler.ctr = ControllerUser{
		usecase: userUsecase,
	}

	userRouter := router.Group("/user")
	userRouter.GET("/detail/:userId", middleware.Authenticate(), handler.FindById)
	userRouter.POST("/login", handler.Login)
	userRouter.POST("/create", handler.Create)
}

func (handler UserRequestHandler) FindById(ctx *gin.Context) {
	userId, err := strconv.ParseUint(ctx.Param("userId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	response, err := handler.ctr.usecase.FindById(ctx, uint(userId))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (handler UserRequestHandler) Login(ctx *gin.Context) {
	var userLogin UserLogin
	err := ctx.BindJSON(&userLogin)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	response, err := handler.ctr.Login(ctx, userLogin)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (handler UserRequestHandler) Create(ctx *gin.Context) {
	var user entity.User
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	response, err := handler.ctr.Create(ctx, user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
