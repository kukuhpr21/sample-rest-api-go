package middleware

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
)

type bodyLogWriter struct {
    gin.ResponseWriter
    body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
    w.body.Write(b)
    return w.ResponseWriter.Write(b)
}

func ResponseLogger(ctx *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
    ctx.Writer = blw
	ctx.Next()
	glg.Log("Response")
	glg.Log("Status     : ", blw.ResponseWriter.Status())
	glg.Log("Body       : ", blw.body.String())
}