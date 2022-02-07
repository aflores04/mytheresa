package main

import (
	"github.com/aflores04/mytheresa/products"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	r := gin.Default()

	products.InitRoutes(r)

	log.Println("products api listening on port 8080")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println(err)
	}

}
