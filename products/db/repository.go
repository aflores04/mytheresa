package db

import (
	"github.com/aflores04/mytheresa/products/config"
	"github.com/aflores04/mytheresa/products/domain"
	"github.com/aflores04/mytheresa/products/request"
)

type ProductRepository interface {
	GetProducts(req *request.GetProductsRequest) ([]*domain.Product, error)
}

type ProductRepositoryImpl struct {
	DbConnection config.DBConnection
}

func NewProductRepository(dbConnection config.DBConnection) ProductRepository {
	return &ProductRepositoryImpl{
		DbConnection: dbConnection,
	}
}

func (r *ProductRepositoryImpl) GetProducts(req *request.GetProductsRequest) ([]*domain.Product, error) {
	var products []*domain.Product

	conn, err := r.DbConnection.GetDB()
	if err != nil {
		return nil, err
	}

	builder := GetProductsQuery()

	if req.Category != nil {
		builder = FilterByCategoryQuery(builder, *req.Category)
	}

	if req.PriceLessThan != nil {
		builder = FilterByPriceLessThanQuery(builder, *req.PriceLessThan)
	}

	rows, err := builder.RunWith(conn).Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := new(domain.Product)
		price := new(domain.Price)

		_ = rows.Scan(
			&product.Sku,
			&product.Name,
			&product.Category,
			&price.Original,
		)

		price.Final = price.Original
		product.Price = price

		products = append(products, product)
	}

	defer conn.Close()

	return products, nil
}
