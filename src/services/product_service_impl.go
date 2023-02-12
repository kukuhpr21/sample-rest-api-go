package services

import (
	"context"
	"kukuhpr21/sample-rest-api-go/src/models/entities"
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
func (s *ProductServiceImpl) Create(ctx context.Context, request productrequest.Create) (response response.ProductResponse, err error) {

	// validate request
	err = s.Validate.Struct(request)

	if err != nil {
		return response, err
	}
	name := request.Name
	result, err := s.ProductRepository.Save(ctx, name)

	if err != nil {
		return response, err
	}

	response.Id = result.Id
	response.Name = result.Name
	return response, nil
}

// Update implements ProductService
func (s *ProductServiceImpl) Update(ctx context.Context, request productrequest.Update) (response response.ProductResponse, err error) {
	// validate request
	err = s.Validate.Struct(request)

	if err != nil {
		return response, err
	}


	result, err := s.ProductRepository.Update(ctx, entities.ProductEntity{
		Id: request.Id,
		Name: request.Name,
	})

	if err != nil {
		return response, err
	}

	response.Id = result.Id
	response.Name = result.Name
	return response, nil
}
