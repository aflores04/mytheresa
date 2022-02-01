package service

import (
	"github.com/aflores04/mytheresa/products/db"
	"github.com/aflores04/mytheresa/products/request"
	"github.com/aflores04/mytheresa/products/response"
)

type ProductService interface {
	GetProducts(req *request.GetProductsRequest) (*response.GetProductsResponse, error)
}

type ProductServiceImpl struct {
	repository db.ProductsRepository
}

func NewProductService(repository db.ProductsRepository) ProductService {
	return &ProductServiceImpl{
		repository: repository,
	}
}

func (s ProductServiceImpl) GetProducts(req *request.GetProductsRequest) (*response.GetProductsResponse, error) {
	products, err := s.repository.GetProducts(req)
	if err != nil {
		return nil, err
	}

	return &response.GetProductsResponse{
		Products: products,
	}, nil
}
