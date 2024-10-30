package user

import (
	"context"
	"errors"
	"exclusive-web/web/dto"
	"exclusive-web/web/entity"
	"exclusive-web/web/repository"
	bcryptpassword "exclusive-web/web/utils/bcrypt"
	jwttoken "exclusive-web/web/utils/bcrypt/jwt-token"
	"time"

	"github.com/google/uuid"
)

type UserUsecase struct {
	repository repository.UserInterfaceRepository
}

type UserUsecaseInterface interface {
	FindById(ctx context.Context, userId uint) (dto.BaseResponse, error)
	Login(ctx context.Context, userLogin UserLogin) (dto.BaseResponse, error)
	Create(ctx context.Context, user entity.User) (dto.BaseResponse, error)
}

func (uc UserUsecase) FindById(ctx context.Context, userId uint) (dto.BaseResponse, error) {

	user, err := uc.repository.FindById(ctx, userId)

	if err != nil {
		return dto.DefaultErrorBaseResponseWithMessage(err)
	}

	return dto.BaseResponse{
		Data:         user,
		Success:      true,
		MessageTitle: "Succeeded",
		Message:      "Get user detail Succeeded",
	}, nil
}

func (uc UserUsecase) Login(ctx context.Context, userlogin UserLogin) (dto.BaseResponse, error) {
	user, err := uc.repository.FindByUsername(ctx, userlogin.Username)

	if err != nil {
		return dto.DefaultErrorBaseResponseWithMessage(err)
	}

	isPasswordCorrect, _ := bcryptpassword.CheckPasswordHash(userlogin.Password, user.Password)

	if !isPasswordCorrect {
		return dto.DefaultErrorBaseResponseWithMessage(errors.New("invalid Password"))
	}

	token, err := jwttoken.GenerateToken(user.Username, user.UserId)

	if err != nil {
		return dto.DefaultErrorBaseResponseWithMessage(err)
	}

	data := LoginResponse{}
	data.Token = token

	return dto.BaseResponse{
		Data:         data,
		Success:      true,
		MessageTitle: "Succeeded",
		Message:      "Login Succeeded",
	}, nil
}

func (uc UserUsecase) Create(ctx context.Context, user entity.User) (dto.BaseResponse, error) {

	HashPassword, _ := bcryptpassword.HashPassword(user.Password)
	userId := uuid.New()
	newUser := &entity.User{
		Username:     user.Username,
		Password:     HashPassword,
		Name:         user.Name,
		Email:        user.Email,
		Phone_Number: user.Phone_Number,
		Address:      user.Address,
		UserId:       userId.String(),
		Created_at:   time.Now(),
		Updated_at:   time.Now(),
	}

	user, err := uc.repository.Create(ctx, *newUser)
	if err != nil {
		return dto.DefaultErrorBaseResponseWithMessage(err)
	}

	return dto.BaseResponse{
		Data:         user,
		Success:      true,
		MessageTitle: "Succeeded",
		Message:      "Create new user Succeeded",
	}, nil
}
