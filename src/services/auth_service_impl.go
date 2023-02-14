package services

import (
	"context"
	"errors"
	"kukuhpr21/sample-rest-api-go/src/models/request/authrequest"
	"kukuhpr21/sample-rest-api-go/src/repositories"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
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
func (s *AuthServiceImpl) Login(ctx context.Context, request authrequest.Login) (interface{}, error) {
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
		return "", errors.New("Invalid email or password")
	}

	accessToken, err := createToken(user.Id, user.Name, user.Email, true)

	if err != nil {
		return "", err
	}

	refreshToken, err := createToken(user.Id, user.Name, user.Email, false)

	if err != nil {
		return "", err
	}

	return map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

func createToken(id, name, email string, isAccess bool) (string, error) {

	type MyClaims struct {
		jwt.StandardClaims
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	timeExpired := os.Getenv("JWT_EXPIRESIN_ACCESS")

	if !isAccess {
		timeExpired = os.Getenv("JWT_EXPIRESIN_REFRESH")
	}

	expiredIn, err := strconv.Atoi(timeExpired)

	if err != nil {
		return "", err
	}

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    os.Getenv("APP_NAME"),
			ExpiresAt: time.Now().Add(time.Duration(expiredIn) * time.Hour).Unix(),
		},
		Id:    id,
		Name:  name,
		Email: email,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
