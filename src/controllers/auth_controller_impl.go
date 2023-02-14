package controllers

import (
	"kukuhpr21/sample-rest-api-go/src/helper"
	"kukuhpr21/sample-rest-api-go/src/models/request/authrequest"
	"kukuhpr21/sample-rest-api-go/src/models/response"
	"kukuhpr21/sample-rest-api-go/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthControllerImpl struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

// Login implements AuthController
func (c *AuthControllerImpl) Login(ctx *gin.Context) {
	var payload *authrequest.Login

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}

	token, err := c.AuthService.Login(ctx.Request.Context(), *payload)

	if err != nil {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}

	helper.SendResponseClient(ctx, response.Client{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   token,
	})
	return
}

// RefreshToken implements AuthController
func (c *AuthControllerImpl) RefreshToken(ctx *gin.Context) {
	refreshToken := ctx.Request.Header.Get("refresh-token")
	token, err := c.AuthService.RefreshToken(ctx, refreshToken)

	if err != nil {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}

	helper.SendResponseClient(ctx, response.Client{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   token,
	})
	return
}
