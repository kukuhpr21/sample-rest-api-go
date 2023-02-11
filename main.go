package main

import (
	"kukuhpr21/sample-rest-api-go/src/config"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/kpango/glg"
)

var (
	server *gin.Engine
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	setupLogger()
	
	glg.Log("======================================SERVICE======================================")
	glg.Log("Name     : " + os.Getenv("APP_NAME"))
	glg.Log("Version  : " + os.Getenv("APP_VERSION"))
	glg.Log("Port     : " + os.Getenv("APP_PORT"))

	// init database
	db, err := config.NewDB(config.DatabaseConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	})

	if err != nil {
		glg.Log("Database : Not Connect [%s]", err.Error())
	} else {
		glg.Log("Database : Connected")
	}

	// init validator
	validate := validator.New()

	currentTime := time.Now()
	glg.Log("Date     : " + currentTime.String())
	glg.Log("======================================SERVICE======================================")

	config.SetupService(config.ServiceConfig{
		Url: os.Getenv("APP_URL"),
		Port: os.Getenv("APP_PORT"),
		Db: db,
		Validate: validate,
	})
}

func setupLogger() {
	logActive := os.Getenv("APP_LOG")

	if logActive == "true" {
		// log := glg.FileWriter("../../log/application.log", 0666)
		log := glg.FileWriter("log/application.log", 0666)
		glg.Get().
			SetMode(glg.BOTH).
			AddLevelWriter(glg.LOG, log).
			AddLevelWriter(glg.DEBG, log).
			AddLevelWriter(glg.INFO, log)
		
		logEr := glg.FileWriter("log/application.err", 0666)
		glg.Get().
			SetMode(glg.BOTH).
			AddLevelWriter(glg.ERR, logEr).
			AddLevelWriter(glg.WARN, logEr)
	}
}
