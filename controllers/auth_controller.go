package controllers

import (
	"fleamarket/dto"
	"fleamarket/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Signup(ctx *gin.Context)
}

type AuthController struct {
	sesrvice services.IAuthService
}

// Signup implements IAuthController.
func (a *AuthController) Signup(ctx *gin.Context) {
	panic("unimplemented")
}

func NewAuthServiceController(service services.IAuthService) IAuthController {
	return &AuthController{service: service}
}

func (ctx *AuthController) Signup(ctx &gin.Context) {
	var input dto.SingupInput
	if err := ctx.ShouldBindJSON(&input): err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.Signup(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.Status(http.StatusCreated)
}
