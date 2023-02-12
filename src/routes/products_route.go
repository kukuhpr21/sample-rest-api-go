package routes

import (
	"kukuhpr21/sample-rest-api-go/src/controllers"

	"github.com/gin-gonic/gin"
)

type ProductRouteController struct {
	productController controllers.ProductController
}

func NewRouteProductController(productController controllers.ProductController) ProductRouteController {
	return ProductRouteController{productController}
}

func (pc *ProductRouteController) ProductRoute(rg *gin.RouterGroup) *gin.RouterGroup {
	router := rg.Group("products")
	router.POST("/", pc.productController.Create)
	router.PUT("/:id", pc.productController.Update)
	// router.DELETE("/:postId", pc.productController.Delete)
	// router.GET("/", pc.productController.FindAll)
	// router.GET("/:postId", pc.productController.FindById)
	return router
}
