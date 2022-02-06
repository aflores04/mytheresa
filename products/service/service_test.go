package service

import (
	"github.com/AlekSi/pointer"
	mocks "github.com/aflores04/mytheresa/mocks/products"
	"github.com/aflores04/mytheresa/products/domain"
	"github.com/aflores04/mytheresa/products/request"
	"github.com/aflores04/mytheresa/products/response"
	"github.com/aflores04/mytheresa/products/utils/discounter"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductService_GetProductsSuccessfully(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var (
		expectedProducts = []*domain.Product{
			&domain.Product{
				Sku:      pointer.ToString("sku"),
				Name:     nil,
				Category: pointer.ToString("my_category"),
				Price: &domain.Price{
					Original: pointer.ToInt64(10000),
					Final:    nil,
					Discount: nil,
					Currency: nil,
				},
			},
		}

		expectedResponse = &response.GetProductsResponse{
			Products: expectedProducts,
		}

		req = request.GetProductsRequest{
			Category:      nil,
			PriceLessThan: nil,
		}
	)

	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().GetProducts(&req).Return(expectedProducts, nil)

	productService := NewProductService(mockProductRepository)
	actualResponse, err := productService.GetProducts(&req)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestProductService_GetProductsWithNilRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)

	productService := NewProductService(mockProductRepository)
	actualResponse, err := productService.GetProducts(nil)

	assert.Nil(t, actualResponse)
	assert.Equal(t, "request cannot be nil", err.Error())
}

func TestProductService_GetProductsWithDiscountByCategory(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var (
		req = request.GetProductsRequest{
			Category:      nil,
			PriceLessThan: nil,
		}

		expectedResponse = &response.GetProductsResponse{
			Products: []*domain.Product{
				{
					Sku:      pointer.ToString("sku"),
					Name:     nil,
					Category: pointer.ToString("boots"),
					Price: &domain.Price{
						Original: pointer.ToInt64(10000),
						Final:    pointer.ToInt64(7000),
						Discount: pointer.ToString("30%"),
						Currency: pointer.ToString(discounter.Currency),
					},
				},
			},
		}
	)

	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().GetProducts(&req).Return([]*domain.Product{
		{
			Sku:      pointer.ToString("sku"),
			Name:     nil,
			Category: pointer.ToString("boots"),
			Price: &domain.Price{
				Original: pointer.ToInt64(10000),
				Final:    nil,
				Discount: nil,
				Currency: nil,
			},
		},
	}, nil)

	productService := NewProductService(mockProductRepository)
	actualResponse, _ := productService.GetProducts(&req)

	assert.Equal(t, expectedResponse, actualResponse)
}

func TestProductService_GetProductsWithDiscountBySKU(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var (
		req = request.GetProductsRequest{
			Category:      nil,
			PriceLessThan: nil,
		}

		expectedResponse = &response.GetProductsResponse{
			Products: []*domain.Product{
				{
					Sku:      pointer.ToString(discounter.SKUWithDiscount),
					Name:     nil,
					Category: pointer.ToString("my-category"),
					Price: &domain.Price{
						Original: pointer.ToInt64(10000),
						Final:    pointer.ToInt64(8500),
						Discount: pointer.ToString("15%"),
						Currency: pointer.ToString(discounter.Currency),
					},
				},
			},
		}
	)

	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().GetProducts(&req).Return([]*domain.Product{
		{
			Sku:      pointer.ToString(discounter.SKUWithDiscount),
			Name:     nil,
			Category: pointer.ToString("my-category"),
			Price: &domain.Price{
				Original: pointer.ToInt64(10000),
				Final:    nil,
				Discount: nil,
				Currency: nil,
			},
		},
	}, nil)

	productService := NewProductService(mockProductRepository)
	actualResponse, _ := productService.GetProducts(&req)

	assert.Equal(t, expectedResponse, actualResponse)
}

func TestProductService_GetProductsWithAllDiscounts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var (
		req = request.GetProductsRequest{
			Category:      nil,
			PriceLessThan: nil,
		}

		expectedResponse = &response.GetProductsResponse{
			Products: []*domain.Product{
				{
					Sku:      pointer.ToString(discounter.SKUWithDiscount),
					Name:     nil,
					Category: pointer.ToString(discounter.CategoryWithDiscount),
					Price: &domain.Price{
						Original: pointer.ToInt64(10000),
						Final:    pointer.ToInt64(5500),
						Discount: pointer.ToString("45%"),
						Currency: pointer.ToString(discounter.Currency),
					},
				},
			},
		}
	)

	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().GetProducts(&req).Return([]*domain.Product{
		{
			Sku:      pointer.ToString(discounter.SKUWithDiscount),
			Name:     nil,
			Category: pointer.ToString(discounter.CategoryWithDiscount),
			Price: &domain.Price{
				Original: pointer.ToInt64(10000),
				Final:    pointer.ToInt64(10000),
				Discount: nil,
				Currency: nil,
			},
		},
	}, nil)

	productService := NewProductService(mockProductRepository)
	actualResponse, _ := productService.GetProducts(&req)

	assert.Equal(t, expectedResponse, actualResponse)
}
