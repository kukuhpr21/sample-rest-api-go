package repositories

import (
	"context"
	"database/sql"
	"errors"
	"kukuhpr21/sample-rest-api-go/src/models/entities"
)

var tUser string = "users"

type UserRepositoryImpl struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) UserRepository {
	return &UserRepositoryImpl{conn: conn}
}

// FindByEmail implements UserRepository
func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (user entities.UserEntity, err error) {
	tx, err := r.conn.Begin()

	if err != nil {
		return user, err
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

	SQL := "SELECT id, id_detail_user, name, email, password, created_at, updated_at FROM products WHERE email = ?"

	// Query
	row := tx.QueryRowContext(ctx, SQL, email)

	err = row.Scan(&user.Id, &user.IdDetailUser, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("User Not Found")
		}
		return user, nil
	}
	return user, nil
}
