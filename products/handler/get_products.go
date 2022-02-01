package handler

import (
	"github.com/aflores04/mytheresa/products/request"
	"github.com/gin-gonic/gin"
	"log"
)

func (h ProductsHandlerImpl) GetProducts(ctx *gin.Context) {
	getProductsRequest := &request.GetProductsRequest{}

	products, err := h.service.GetProducts(getProductsRequest)
	if err != nil {
		log.Println(err)
	}

	ctx.JSON(200, products)
}
