package helper

import (
	"errors"
	"kukuhpr21/sample-rest-api-go/src/models/response"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SendResponseClient(ctx *gin.Context, response response.Client) {
	ctx.JSON(response.Code, response)
}

func IsValidToken(tokenRequest string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenRequest, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, errors.New("Signing method invalid")
		}

		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}
