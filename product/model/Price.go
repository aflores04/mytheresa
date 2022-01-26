package model

type Price struct {
	Original	int64	`json:"original"`
	Final		int64	`json:"final"`
	Discount	string	`json:"discount_percentage"`
	Currency	string	`json:"currency"`
}
