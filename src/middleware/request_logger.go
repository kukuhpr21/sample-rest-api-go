package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
)

func RequestLogger(ctx *gin.Context) {
	request := ctx.Request
	glg.Log("\n\n\n")
	glg.Log("Request")
	glg.Log("Path       : ", request.URL.Path)
	glg.Log("Method     : ", request.Method)
	glg.Log("User Agent : ", request.UserAgent())
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		glg.Log("Body       : error parsing ", err)
	}
	
	buffer := new(bytes.Buffer)

	if err := json.Compact(buffer, body); err != nil {
		glg.Log("Body       : error json compact ", err)
	}
	glg.Log("Body       : ", buffer)

}