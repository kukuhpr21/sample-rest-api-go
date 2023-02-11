package services

import (
	"context"
	"kukuhpr21/sample-rest-api-go/src/models/request/productrequest"
	"kukuhpr21/sample-rest-api-go/src/models/response"
	"kukuhpr21/sample-rest-api-go/src/repositories"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repositories.ProductRepository
	Validate          *validator.Validate
}

func NewProductService(productRepository repositories.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		Validate:          validate,
	}
}

// Create implements ProductService
func (service *ProductServiceImpl) Create(ctx context.Context, request productrequest.Create) (response response.ProductResponse, err error) {

	// validate request
	err = service.Validate.Struct(request)

	if err != nil {
		return response, err
	}
	name := request.Name
	productEntity, err := service.ProductRepository.Save(ctx, name)

	if err != nil {
		return response, err
	}

	response.Id = productEntity.Id
	response.Name = productEntity.Name
	return response, nil
}
