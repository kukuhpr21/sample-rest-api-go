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
func (r *ProductRepositoryImpl) Save(ctx context.Context, name string) (product entities.ProductEntity, err error) {

	tx, err := r.conn.Begin()

	if err != nil {
		return product, err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
		} else if err != nil {
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
func (r *ProductRepositoryImpl) Update(ctx context.Context, product entities.ProductEntity) (entities.ProductEntity, error) {
	tx, errTx := r.conn.Begin()

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

	SQL := "UPDATE " + tProduct + " SET name = ? WHERE id = ?"

	// Exec
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Id)

	if err != nil {
		return product, err
	}
	return product, nil
}

// Delete implements ProductRepository
func (r *ProductRepositoryImpl) Delete(ctx context.Context, id int) (err error) {
	tx, err := r.conn.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	SQL := "DELETE FROM " + tProduct + " WHERE id = ?"

	// Exec
	_, err = tx.ExecContext(ctx, SQL, id)

	if err != nil {
		return err
	}
	return nil

}

// FindAll implements ProductRepository
func (r *ProductRepositoryImpl) FindAll(ctx context.Context) (products []entities.ProductEntity, err error) {
	tx, err := r.conn.Begin()

	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	SQL := "SELECT id, name FROM products"

	// Query
	rows, err := tx.QueryContext(ctx, SQL)
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		product := entities.ProductEntity{} // buat empty object categories
		rows.Scan(&product.Id, &product.Name)
		
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// FindById implements ProductRepository
func (r *ProductRepositoryImpl) FindById(ctx context.Context, id int) (product entities.ProductEntity, err error) {
	tx, err := r.conn.Begin()

	if err != nil {
		return product, err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	SQL := "SELECT id, name FROM products WHERE id = ?"

	// Query
	rows, err := tx.QueryContext(ctx, SQL, id)

	if err != nil {
		return product, err
	}

	if rows.Next() {
		rows.Scan(&product.Id, &product.Name)
		return product, nil
	}
	return product, nil
}
