package products

import (
	"github.com/aflores04/mytheresa/products/config"
	"github.com/aflores04/mytheresa/products/db"
	"github.com/aflores04/mytheresa/products/handler"
	"github.com/aflores04/mytheresa/products/service"
	"github.com/gin-gonic/gin"
	"os"
)

func InitRoutes(r *gin.Engine) {
	dbConnection := config.NewDBConnection(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	productsRepository := db.NewProductRepository(dbConnection)
	productsService := service.NewProductService(productsRepository)
	productsHandler := handler.NewProductsHandler(productsService)

	r.GET("/products", productsHandler.GetProducts)
}
