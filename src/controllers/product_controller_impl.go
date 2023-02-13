package controllers

import (
	"kukuhpr21/sample-rest-api-go/src/helper"
	"kukuhpr21/sample-rest-api-go/src/models/request/productrequest"
	"kukuhpr21/sample-rest-api-go/src/models/response"
	"kukuhpr21/sample-rest-api-go/src/services"
	"net/http"
	"strconv"

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
func (c *ProductControllerImpl) Create(ctx *gin.Context) {
	var payload *productrequest.Create

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}

	data, err := c.ProductService.Create(ctx.Request.Context(), *payload)

	if err != nil {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Data:   err.Error(),
		})
		return
	} else {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusCreated,
			Status: http.StatusText(http.StatusCreated),
			Data:   data,
		})
		return
	}
}

// Update implements ProductController
func (c *ProductControllerImpl) Update(ctx *gin.Context) {
	var payload *productrequest.Update

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}

	productId := ctx.Params.ByName("id")
	id, err := strconv.Atoi(productId)

	if err != nil {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}
	payload.Id = id
	data, err := c.ProductService.Update(ctx.Request.Context(), *payload)

	if err != nil {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Data:   err.Error(),
		})
		return
	} else {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   data,
		})
		return
	}
}

// Delete implements ProductController
func (c *ProductControllerImpl) Delete(ctx *gin.Context) {
	productId := ctx.Params.ByName("id")
	id, err := strconv.Atoi(productId)

	if err != nil {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}

	err = c.ProductService.Delete(ctx.Request.Context(), id)

	if err != nil {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Data:   err.Error(),
		})
		return
	} else {
		type data struct {
			Id int
		}
		mData := data{
			Id: id,
		}
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   mData,
		})
		return
	}
}

// FindAll implements ProductController
func (c *ProductControllerImpl) FindAll(ctx *gin.Context) {
	datas, err := c.ProductService.FindAll(ctx.Request.Context())

	if err != nil {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Data:   err.Error(),
		})
		return
	} else {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   datas,
		})
		return
	}
}

// FindById implements ProductController
func (c *ProductControllerImpl) FindById(ctx *gin.Context) {
	productId := ctx.Params.ByName("id")
	id, err := strconv.Atoi(productId)

	if err != nil {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}

	data, err := c.ProductService.FindById(ctx.Request.Context(), id)

	if err != nil {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Data:   err.Error(),
		})
		return
	}

	if data.Name == "" {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Data: "",
		})
		return
	} else {
		helper.SendResponseClient(ctx, response.Client{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data: data,
		})
		return
	}

}
