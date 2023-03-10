package services

import (
	"context"
	"kukuhpr21/sample-rest-api-go/src/models/request/productrequest"
	"kukuhpr21/sample-rest-api-go/src/models/response"
)

type ProductService interface {
	Create(ctx context.Context, request productrequest.Create) (response.ProductResponse, error)
	Update(ctx context.Context, request productrequest.Update) (response.ProductResponse, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]response.ProductResponse, error)
	FindById(ctx context.Context, id int) (response.ProductResponse, error)
}
