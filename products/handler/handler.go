package handler

import (
	"github.com/aflores04/mytheresa/products/service"
	"github.com/gin-gonic/gin"
)

type ProductsHandler interface {
	GetProducts(ctx *gin.Context)
}

type ProductsHandlerImpl struct {
	service service.ProductService
}

func NewProductsHandler(service service.ProductService) ProductsHandler {
	return &ProductsHandlerImpl{
		service: service,
	}
}