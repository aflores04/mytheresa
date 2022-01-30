package domain

type Product struct {
	Sku			string	`json:"sku"`
	Name		string	`json:"name"`
	Category	string	`json:"category"`
	Price		Price	`json:"price"`
}
