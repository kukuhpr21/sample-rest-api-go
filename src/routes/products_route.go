package routes

import (
	"kukuhpr21/sample-rest-api-go/src/controllers"
	"kukuhpr21/sample-rest-api-go/src/middleware"

	"github.com/gin-gonic/gin"
)

type ProductRouteController struct {
	productController controllers.ProductController
}

func NewProductRouteController(productController controllers.ProductController) ProductRouteController {
	return ProductRouteController{productController}
}

func (pc *ProductRouteController) ProductRoute(rg *gin.RouterGroup) *gin.RouterGroup {
	router := rg.Group("products")
	router.Use(middleware.VerifyToken())
	router.POST("/", pc.productController.Create)
	router.PUT("/:id", pc.productController.Update)
	router.DELETE("/:id", pc.productController.Delete)
	router.GET("/", pc.productController.FindAll)
	router.GET("/:id", pc.productController.FindById)
	return router
}
