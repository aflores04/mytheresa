package service

import (
	"github.com/aflores04/mytheresa/products/db"
	"github.com/aflores04/mytheresa/products/errors"
	"github.com/aflores04/mytheresa/products/request"
	"github.com/aflores04/mytheresa/products/response"
	"github.com/aflores04/mytheresa/products/utils/discounter"
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
		priceDiscounter := discounter.NewDiscounter(*product.Price.Original)

		if *product.Category == discounter.CategoryWithDiscount {
			priceDiscounter.DiscountAmount += discounter.CategoryDiscountAmount
		}

		if *product.Sku == discounter.SKUWithDiscount {
			priceDiscounter.DiscountAmount += discounter.SKUDiscountAmount
		}

		product.Price = priceDiscounter.FillPriceWithDiscount()
	}

	return &response.GetProductsResponse{
		Products: products,
	}, nil
}
