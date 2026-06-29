package http

import (
	"github.com/gin-gonic/gin"

	"rent-manager-backend/internal/core/domain"
	userUseCase "rent-manager-backend/internal/core/usecase/user"
)

type AuthHandler struct {
	loginUseCase *userUseCase.LoginUseCase
}

func NewAuthHandler(loginUseCase *userUseCase.LoginUseCase) *AuthHandler {
	return &AuthHandler{
		loginUseCase: loginUseCase,
	}
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req loginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	token, err := h.loginUseCase.Login(ctx, &domain.Login{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, loginResponse{
		Token: token,
	})
}
