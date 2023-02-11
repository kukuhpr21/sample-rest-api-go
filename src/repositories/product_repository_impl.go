package repositories

import (
	"context"
	"database/sql"
	"kukuhpr21/sample-rest-api-go/src/models/entities"
)

var tProduct string = "products"

type ProductRepositoryImpl struct {
	conn *sql.DB
}

func NewProductRepository(conn *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{conn: conn}
}

// Save implements ProductRepository
func (pr *ProductRepositoryImpl) Save(ctx context.Context, name string) (entities.ProductEntity, error) {
	product := entities.ProductEntity{}

	tx, errTx := pr.conn.Begin()

	if errTx != nil {
		return product, errTx
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
		} else if errTx != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	SQL := "INSERT INTO " + tProduct + " (name) VALUES (?)"

	// Exec
	result, err := tx.ExecContext(ctx, SQL, name)

	if err != nil {
		return product, err
	}

	// Get LastInsertId
	id, err := result.LastInsertId()

	if err != nil {
		return product, err
	}

	product.Id = int(id)
	product.Name = name

	// Return data
	return product, err
}

// Update implements ProductRepository
func (*ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, id int, product entities.ProductEntity) (entities.ProductEntity, error) {
	panic("unimplemented")
}

// Delete implements ProductRepository
func (*ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	panic("unimplemented")
}

// FindAll implements ProductRepository
func (*ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]entities.ProductEntity, error) {
	panic("unimplemented")
}

// FindById implements ProductRepository
func (*ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entities.ProductEntity, error) {
	panic("unimplemented")
}
