package service

import (
	"fmt"
	"github.com/AlekSi/pointer"
	"github.com/aflores04/mytheresa/products/db"
	"github.com/aflores04/mytheresa/products/errors"
	"github.com/aflores04/mytheresa/products/request"
	"github.com/aflores04/mytheresa/products/response"
	"math"
)

const (
	Currency               string = "EUR"
	CategoryWithDiscount   string = "boots"
	CategoryDiscountAmount int64  = 30
	SKUDiscount            int64  = 15
)

type ProductService interface {
	GetProducts(req *request.GetProductsRequest) (*response.GetProductsResponse, error)
}

type ProductServiceImpl struct {
	repository db.ProductRepository
}

func NewProductService(repository db.ProductRepository) ProductService {
	return &ProductServiceImpl{
		repository: repository,
	}
}

func (s ProductServiceImpl) GetProducts(req *request.GetProductsRequest) (*response.GetProductsResponse, error) {
	if req == nil {
		return nil, errors.NewError("request cannot be nil", nil)
	}

	products, err := s.repository.GetProducts(req)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		if *product.Category == CategoryWithDiscount {
			product.Price.Discount = pointer.ToString("%" + fmt.Sprintf("%v", CategoryDiscountAmount))

			priceDiscountApplied, err := applyDiscount(*product.Price.Original, CategoryDiscountAmount)
			if err != nil {
				return nil, errors.NewError("error when applying discount: %s", err)
			}

			product.Price.Final = priceDiscountApplied
		}
	}

	return &response.GetProductsResponse{
		Products: products,
	}, nil
}

func applyDiscount(number int64, discount int64) (*int64, error) {
	var (
		priceWithDecimals float64
		priceDiscount     float64
	)

	if number == 0 {
		return nil, errors.NewError("price cannot be 0", nil)
	}

	// format number with decimals
	priceWithDecimals = float64(number / 100)

	// calculate discount amount
	priceDiscount = (priceWithDecimals * float64(discount)) / 100

	// round to nearest from more than two decimals to only two ... i.e 0.543 -> 0.54
	round := math.Round(priceDiscount*100) / 100

	// apply discount to price and convert to integer
	finalPrice := int64((priceWithDecimals - round) * 100)

	return &finalPrice, nil
}
