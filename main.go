package main

import (
	"github.com/aflores04/mytheresa/products/config"
	"github.com/aflores04/mytheresa/products/db"
	"github.com/aflores04/mytheresa/products/handler"
	"github.com/aflores04/mytheresa/products/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	dbConnection := config.NewDBConnection(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	productsRepository := db.NewProductRepository(dbConnection)
	productsService := service.NewProductService(productsRepository)
	productsHandler := handler.NewProductsHandler(productsService)

	r := gin.Default()

	r.GET("/products", productsHandler.GetProducts)

	log.Println("products api listening on port 8080")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println(err)
	}

}
