package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"

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

func HTTPLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request logger
		t := time.Now()
		request := c.Request

		glg.Log("\n\n\n")
		glg.Log("Request")
		glg.Log("Path       : ", request.URL.Path)
		glg.Log("Method     : ", request.Method)
		glg.Log("User Agent : ", request.UserAgent())
		
		body, err := ioutil.ReadAll(request.Body)
	
		if err != nil {
			glg.Log("Body       :  error parsing ", err.Error())
		}

		if len(body) > 0 {

			buffer := new(bytes.Buffer)
			if err := json.Compact(buffer, body); err != nil {
				glg.Log("Body       :  error json compact ", err)
			} else {
				glg.Log("Body       : ", buffer)
			}
		} else {
			glg.Log("Body       :  empty")
		}

		c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
    	c.Writer = blw

		c.Next()
	
		// response logger
		duration := time.Since(t)
		
		glg.Log("")
		glg.Log("Response")
		glg.Log("Status     : ", blw.ResponseWriter.Status())
		glg.Log("Body       : ", blw.body.String())
		glg.Log("Duration   : ", duration)
	}
}