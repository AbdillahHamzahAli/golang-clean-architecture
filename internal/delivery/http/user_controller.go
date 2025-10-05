package http

import (
	"net/http"

	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/domain/dto"
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/shared/response"
	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UseCase usecase.AuthUsecase
}

func NewUserController(useCase usecase.AuthUsecase) *UserController {
	return &UserController{
		UseCase: useCase,
	}
}

func (uc *UserController) Register(ctx *gin.Context) {
	var user dto.RegisterUserRequest
	err := ctx.ShouldBindBodyWithJSON(&user)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	resp, err := uc.UseCase.Register(ctx.Request.Context(), user)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "register failed", err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, "register success", resp)
}

func (uc *UserController) Login(ctx *gin.Context) {
	var user dto.LoginUserRequest
	err := ctx.ShouldBindBodyWithJSON(&user)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	resp, err := uc.UseCase.Login(ctx.Request.Context(), user)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "login failed", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "login success", resp)
}
