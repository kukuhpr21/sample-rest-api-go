package middleware

import (
	"kukuhpr21/sample-rest-api-go/src/helper"
	"kukuhpr21/sample-rest-api-go/src/models/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")

		if strings.HasPrefix(authorization, "Bearer") {
			tokenRequest := strings.Replace(authorization, "Bearer ", "", 1)

			_, err := helper.IsValidToken(tokenRequest)

			if err != nil {
				helper.SendResponseClient(c, response.Client{
					Code:   http.StatusBadRequest,
					Status: http.StatusText(http.StatusBadRequest),
					Data:   "Invalid token",
				})
				c.Abort()
			} else {
				c.Next()
			}
		} else {
			helper.SendResponseClient(c, response.Client{
				Code:   http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Data:   "Invalid token",
			})
			c.Abort()
		}
	}
}
