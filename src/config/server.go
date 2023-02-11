package config

import (
	"database/sql"
	"io/ioutil"
	"kukuhpr21/sample-rest-api-go/src/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kpango/glg"
)

type ServiceConfig struct {
	Url   string
	Port string
	Db *sql.DB
	Validate *validator.Validate
}

func SetupService(c ServiceConfig) {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	server := gin.Default()
	server.Use(gin.Recovery())
	server.Use(middleware.RequestLogger)
	server.Use(middleware.ResponseLogger)
	router := server.Group("/v1")
	SetupLayer(LayerConfig {
		R: router,
		Db: c.Db,
		V: c.Validate,
	})
	err := server.Run(c.Url + ":" + c.Port)

	if err != nil {
		glg.Log("Server   : Not Connect [%s]", err.Error())
	}
}