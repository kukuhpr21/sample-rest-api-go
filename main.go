package main

import (
	"fmt"
	"kukuhpr21/sample-rest-api-go/src/config"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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
	fmt.Println("======================================SERVICE======================================")
	fmt.Println("Name     : " + os.Getenv("APP_NAME"))
	fmt.Println("Version  : " + os.Getenv("APP_VERSION"))
	fmt.Println("Port     : " + os.Getenv("APP_PORT"))

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
		fmt.Printf("Database : Not Connect [%s]", err.Error())
	} else {
		fmt.Println("Database : Connected")
	}

	// init validator
	validate := validator.New()

	currentTime := time.Now()
	fmt.Println("Date     : " + currentTime.String())
	fmt.Println("======================================SERVICE======================================")

	config.SetupService(config.SetupServiceConfig{
		Url: os.Getenv("APP_URL"),
		Port: os.Getenv("APP_PORT"),
		Db: db,
		Validate: validate,
	})
}
