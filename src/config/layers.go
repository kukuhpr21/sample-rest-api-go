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
	R *gin.RouterGroup
	Db *sql.DB
	V *validator.Validate
}

func SetupLayer(c LayerConfig) {
	productRoute := productLayers(c.Db, c.V)

	productRoute.ProductRoute(c.R)

}

func productLayers(db *sql.DB, v *validator.Validate) routes.ProductRouteController {
	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository, v)
	productController := controllers.NewProductController(productService)
	productRoute := routes.NewRouteProductController(productController)
	return productRoute
}