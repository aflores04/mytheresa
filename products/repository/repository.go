package repository

import (
	"github.com/aflores04/mytheresa/products/config"
	"github.com/aflores04/mytheresa/products/domain"
	"github.com/aflores04/mytheresa/products/request"
)

type ProductsRepository interface {
	GetProducts(req *request.GetProductRequest) ([]*domain.Product, error)
}

type ProductRepositoryImpl struct {
	DbConnection config.DBConnection
}

func NewProductRepository(dbConnection config.DBConnection) ProductsRepository {
	return &ProductRepositoryImpl{
		DbConnection: dbConnection,
	}
}

func (r *ProductRepositoryImpl) GetProducts(req *request.GetProductRequest) ([]*domain.Product, error) {
	return []*domain.Product{}, nil
}