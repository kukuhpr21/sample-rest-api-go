package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SetupServiceConfig struct {
	Url   string
	Port string
	Db *sql.DB
	Validate *validator.Validate
}

func SetupService(c SetupServiceConfig) {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	server := gin.Default()
	server.Use(gin.Recovery())

	router := server.Group("/v1")
	RegisterLayers(router, c.Db, c.Validate)
	err := server.Run(c.Url + ":" + c.Port)

	if err != nil {
		fmt.Printf("Server   : Not Connect [%s]", err.Error())
	}
}