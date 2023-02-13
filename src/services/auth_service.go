package services

import (
	"context"
	"kukuhpr21/sample-rest-api-go/src/models/request/authrequest"
)

type AuthService interface {
	Login(ctx context.Context, request authrequest.Login) (string, error)
}
