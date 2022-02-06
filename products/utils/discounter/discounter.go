package discounter

import (
	"github.com/AlekSi/pointer"
	"github.com/aflores04/mytheresa/products/domain"
	"github.com/aflores04/mytheresa/products/helpers"
	"strconv"
)

const (
	Currency               string = "EUR"
	CategoryWithDiscount   string = "boots"
	CategoryDiscountAmount int64  = 30
	SKUWithDiscount        string = "000003"
	SKUDiscountAmount      int64  = 15
)

type Discounter struct {
	DiscountAmount int64
	Price          int64
}

func NewDiscounter(price int64) *Discounter {
	return &Discounter{
		Price: price,
	}
}

func (d *Discounter) Accumulate(amount int64) {
	d.DiscountAmount = d.DiscountAmount + amount
}

func (d Discounter) FillPriceWithDiscount() *domain.Price {
	totalAmountDiscount := strconv.Itoa(int(d.DiscountAmount))

	return &domain.Price{
		Original: pointer.ToInt64(d.Price),
		Final:    pointer.ToInt64(helpers.ApplyDiscount(d.Price, d.DiscountAmount)),
		Discount: pointer.ToString(totalAmountDiscount + "%"),
		Currency: pointer.ToString(Currency),
	}
}
