package response

import "github.com/aflores04/mytheresa/products/domain"

type GetProductsResponse struct {
	Products []*domain.Product `json:"products"`
}
