package services

import (
	"context"
	"kukuhpr21/sample-rest-api-go/src/models/request/productrequest"
	"kukuhpr21/sample-rest-api-go/src/models/response"
)

type ProductService interface {
	Create(ctx context.Context, request productrequest.Create) (response.ProductResponse, error)
	// Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	// Delete(ctx context.Context, categoryId int)
	// FindById(ctx context.Context, categoryId int) web.CategoryResponse
	// FindAll(ctx context.Context) []web.CategoryResponse
}
