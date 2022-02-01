package request

type GetProductsRequest struct {
	Category      *string `json:"category"`
	PriceLessThan *int64  `json:"price_less_than"`
}
