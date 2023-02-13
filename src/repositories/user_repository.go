package repositories

import (
	"context"
	"kukuhpr21/sample-rest-api-go/src/models/entities"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (entities.UserEntity, error)
}
