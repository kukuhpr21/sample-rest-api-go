package controllers

import (
	"kukuhpr21/sample-rest-api-go/src/helper"
	"kukuhpr21/sample-rest-api-go/src/models/request/productrequest"
	"kukuhpr21/sample-rest-api-go/src/models/response"
	"kukuhpr21/sample-rest-api-go/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

// Create implements ProductController
func (p *ProductControllerImpl) Create(ctx *gin.Context) {
	var payload *productrequest.Create

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}

	data, err := p.ProductService.Create(ctx.Request.Context(), *payload)

	if err != nil {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Data:   err.Error(),
		})
	} else {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusCreated,
			Status: http.StatusText(http.StatusCreated),
			Data:   data,
		})
	}

}
