package helpers

func ApplyDiscount(number int64, discount int64) int64 {
	var (
		priceDiscount float64
	)

	// calculate discounter amount
	priceDiscount = float64((number * discount) / 100)

	// apply discounter to price
	finalPrice := float64(number) - priceDiscount

	rounded := int64(finalPrice)

	return rounded
}
