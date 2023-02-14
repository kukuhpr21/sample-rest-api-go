package routes

import (
	"kukuhpr21/sample-rest-api-go/src/controllers"

	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewRouteAuthController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (ac *AuthRouteController) AuthRoute(rg *gin.RouterGroup) *gin.RouterGroup {
	router := rg.Group("auth")
	router.POST("/login", ac.authController.Login)
	router.GET("/refresh-token", ac.authController.RefreshToken)
	return router
}
