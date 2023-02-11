package helper

import (
	"kukuhpr21/sample-rest-api-go/src/models/response"

	"github.com/gin-gonic/gin"
)

func SendResponseClient(ctx *gin.Context, response response.Client) {
	ctx.JSON(response.Code, response)
}
