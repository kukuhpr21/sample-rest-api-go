package main

import (
	"fmt"
	"io/ioutil"
	"kukuhpr21/sample-rest-api-go/src/config"
	"kukuhpr21/sample-rest-api-go/src/controllers"
	"kukuhpr21/sample-rest-api-go/src/repositories"
	"kukuhpr21/sample-rest-api-go/src/routes"
	"kukuhpr21/sample-rest-api-go/src/services"
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

	fmt.Println("")
	fmt.Println("======================================SERVICE======================================")
	fmt.Println("Name : " + os.Getenv("APP_NAME"))
	fmt.Println("Version : " + os.Getenv("APP_VERSION"))
	fmt.Println("Port : " + os.Getenv("APP_PORT"))

	// init database
	db, err := config.NewDB(config.DatabaseConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Username: os.Getenv("DB_USERNAME"),
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

	// init layers
	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository, validate)
	productController := controllers.NewProductController(productService)
	productRoute := routes.NewRouteProductController(productController)

	currentTime := time.Now()
	fmt.Println("Date Time : " + currentTime.String())
	fmt.Println("======================================SERVICE======================================")

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	server := gin.Default()
	server.Use(gin.Recovery())

	router := server.Group("/api")
	productRoute.ProductRoute(router)
	err = server.Run(os.Getenv("APP_URL") + ":" + os.Getenv("APP_PORT"))

	if err != nil {
		fmt.Printf("Server : Not Connect [%s]", err.Error())
	}
}
