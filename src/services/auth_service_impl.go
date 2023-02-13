package services

import (
	"context"
	"errors"
	"kukuhpr21/sample-rest-api-go/src/models/request/authrequest"
	"kukuhpr21/sample-rest-api-go/src/repositories"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/kpango/glg"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	UserRepository repositories.UserRepository
	Validate       *validator.Validate
}

func NewAuthService(userRepository repositories.UserRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

// Login implements AuthService
func (s *AuthServiceImpl) Login(ctx context.Context, request authrequest.Login) (string, error) {
	err := s.Validate.Struct(request)

	if err != nil {
		return "", err
	}
	user, err := s.UserRepository.FindByEmail(ctx, request.Email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		glg.Error(err.Error())
		return "", errors.New("Invalid email or password")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
	})
	return token.Raw, nil
}

// // FindByEmail implements UserService
// func (s *UserServiceImpl) FindByEmail(ctx context.Context, email string) (data response.UserResponse, err error) {
// 	user, err := s.UserRepository.FindByEmail(ctx, email)

// 	if err != nil {
// 		return data, err
// 	}

// 	data.Id = user.Id
// 	data.IdDetailUser = user.IdDetailUser
// 	data.Name = user.Name
// 	data.Email = user.Email
// 	data.CreatedAt = user.CreatedAt
// 	data.UpdatedAt = user.UpdatedAt
// 	return data, nil
// }
