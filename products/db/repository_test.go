package db

import (
	"github.com/AlekSi/pointer"
	"github.com/aflores04/mytheresa/products/config"
	"github.com/aflores04/mytheresa/products/domain"
	"github.com/aflores04/mytheresa/products/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductRepository_GetProducts(t *testing.T) {
	// added connection here because in concordance to the
	connection := config.NewDBConnection(
		"127.0.0.1",
		"5432",
		"products_testing_user",
		"products_testing_password",
		"products_testing_database")

	repository := NewProductRepository(connection)

	tests := []struct {
		name     string
		input    request.GetProductsRequest
		expected []*domain.Product
	}{
		{
			name: "filter by category and price",
			input: request.GetProductsRequest{
				Category:      pointer.ToString("cat1"),
				PriceLessThan: pointer.ToInt64(200),
			},
			expected: []*domain.Product{
				&domain.Product{
					Sku:      pointer.ToString("1"),
					Name:     pointer.ToString("test1"),
					Category: pointer.ToString("cat1"),
					Price: &domain.Price{
						Original: 100,
						Final:    0,
						Discount: "",
						Currency: "",
					},
				},
			},
		},
		{
			name: "filter by category",
			input: request.GetProductsRequest{
				Category: pointer.ToString("cat1"),
			},
			expected: []*domain.Product{
				&domain.Product{
					Sku:      pointer.ToString("1"),
					Name:     pointer.ToString("test1"),
					Category: pointer.ToString("cat1"),
					Price: &domain.Price{
						Original: 100,
						Final:    0,
						Discount: "",
						Currency: "",
					},
				},
			},
		},
		{
			name: "filter by price",
			input: request.GetProductsRequest{
				PriceLessThan: pointer.ToInt64(200),
			},
			expected: []*domain.Product{
				&domain.Product{
					Sku:      pointer.ToString("1"),
					Name:     pointer.ToString("test1"),
					Category: pointer.ToString("cat1"),
					Price: &domain.Price{
						Original: 100,
						Final:    0,
						Discount: "",
						Currency: "",
					},
				},
			},
		},
		{
			name:  "get all products",
			input: request.GetProductsRequest{},
			expected: []*domain.Product{
				&domain.Product{
					Sku:      pointer.ToString("1"),
					Name:     pointer.ToString("test1"),
					Category: pointer.ToString("cat1"),
					Price: &domain.Price{
						Original: 100,
						Final:    0,
						Discount: "",
						Currency: "",
					},
				},
				&domain.Product{
					Sku:      pointer.ToString("2"),
					Name:     pointer.ToString("test2"),
					Category: pointer.ToString("cat2"),
					Price: &domain.Price{
						Original: 200,
						Final:    0,
						Discount: "",
						Currency: "",
					},
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			products, err := repository.GetProducts(&testCase.input)

			assert.Nil(t, err)
			assert.Equal(t, testCase.expected, products)
		})
	}
}
