package config

import (
	"database/sql"
	"kukuhpr21/sample-rest-api-go/src/controllers"
	"kukuhpr21/sample-rest-api-go/src/repositories"
	"kukuhpr21/sample-rest-api-go/src/routes"
	"kukuhpr21/sample-rest-api-go/src/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LayerConfig struct {
	R  *gin.RouterGroup
	Db *sql.DB
	V  *validator.Validate
}

func SetupLayer(c LayerConfig) {
	authRoute := authLayers(c.Db, c.V)
	productRoute := productLayers(c.Db, c.V)

	authRoute.AuthRoute(c.R)
	productRoute.ProductRoute(c.R)

}

func authLayers(db *sql.DB, v *validator.Validate) routes.AuthRouteController {
	repository := repositories.NewUserRepository(db)
	service := services.NewAuthService(repository, v)
	controller := controllers.NewAuthController(service)
	route := routes.NewAuthRouteController(controller)
	return route
}

func productLayers(db *sql.DB, v *validator.Validate) routes.ProductRouteController {
	repository := repositories.NewProductRepository(db)
	service := services.NewProductService(repository, v)
	controller := controllers.NewProductController(service)
	route := routes.NewProductRouteController(controller)
	return route
}
