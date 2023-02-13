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
		Id:   request.Id,
		Name: request.Name,
	})

	if err != nil {
		return response, err
	}

	response.Id = result.Id
	response.Name = result.Name
	return response, nil
}

// Delete implements ProductService
func (s *ProductServiceImpl) Delete(ctx context.Context, id int) error {
	err := s.ProductRepository.Delete(ctx, id)

	if err != nil {
		return err
	}
	return nil
}

// FindAll implements ProductService
func (s *ProductServiceImpl) FindAll(ctx context.Context) (datas []response.ProductResponse, err error) {
	products, err := s.ProductRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(products); i++ {
		data := response.ProductResponse{}
		data.Id = products[i].Id
		data.Name = products[i].Name
		datas = append(datas, data)
	}
	return datas, nil
}

// FindById implements ProductService
func (s *ProductServiceImpl) FindById(ctx context.Context, id int) (data response.ProductResponse, err error) {
	product, err := s.ProductRepository.FindById(ctx, id)

	if err != nil {
		return data, err
	}

	data.Id = product.Id
	data.Name = product.Name
	return data, nil
}
