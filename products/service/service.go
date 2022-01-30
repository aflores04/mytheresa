package service

import (
	"github.com/aflores04/mytheresa/products/repository"
	"github.com/aflores04/mytheresa/products/request"
	"github.com/aflores04/mytheresa/products/response"
)

type ProductService interface {
	GetProducts(req *request.GetProductRequest) (*response.GetProductsResponse, error)
}

type ProductServiceImpl struct {
	repository repository.ProductsRepository
}

func NewProductService(repository repository.ProductsRepository) ProductService {
	return &ProductServiceImpl{
		repository: repository,
	}
}

func (s ProductServiceImpl) GetProducts(req *request.GetProductRequest) (*response.GetProductsResponse, error) {
	products, err := s.repository.GetProducts(req)
	if err != nil {
		return nil, err
	}

	return &response.GetProductsResponse{
		Products: products,
	}, nil
}