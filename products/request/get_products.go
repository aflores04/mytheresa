package request

type GetProductRequest struct {
	Category 		string 	`json:"category"`
	PriceLessThan	int		`json:"price_less_than"`
}
