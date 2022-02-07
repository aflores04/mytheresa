package handler

import (
	"github.com/AlekSi/pointer"
	"github.com/aflores04/mytheresa/products/request"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func (h ProductsHandlerImpl) GetProducts(ctx *gin.Context) {
	getProductsRequest := queryStringToRequest(ctx)

	res, err := h.service.GetProducts(getProductsRequest)
	if err != nil {
		log.Println(err)
	}

	ctx.JSON(200, res)
}

// queryStringToRequest fill request with query strings
func queryStringToRequest(ctx *gin.Context) *request.GetProductsRequest {
	var (
		category *string = nil

		req = request.GetProductsRequest{
			Category:      nil,
			PriceLessThan: nil,
		}
	)

	categoryQuery := ctx.Query("category")
	priceLessThanQuery := ctx.Query("price_less_than")

	if categoryQuery != "" {
		category = pointer.ToString(categoryQuery)
		req.Category = category
	}

	// parse priceLessThan to string, if empty will be 0
	p, _ := strconv.Atoi(priceLessThanQuery)

	req.PriceLessThan = pointer.ToInt64(int64(p))

	// priceLessThan cannot be 0 or empty, in both cases it will be nil
	// we need nils for apply or not filters in repository
	if p == 0 {
		req.PriceLessThan = nil
	}

	return &req
}
