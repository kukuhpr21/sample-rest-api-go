package repositories

import (
	"context"
	"kukuhpr21/sample-rest-api-go/src/models/entities"
)

type ProductRepository interface {
	Save(ctx context.Context, name string) (entities.ProductEntity, error)
	Update(ctx context.Context, product entities.ProductEntity) (entities.ProductEntity, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]entities.ProductEntity, error)
	FindById(ctx context.Context, id int) (entities.ProductEntity, error)
}
