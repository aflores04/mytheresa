package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestProductService_GetProducts(t *testing.T) {
//	mockCtrl := gomock.NewController(t)
//	defer mockCtrl.Finish()
//
//	req := request.GetProductsRequest{
//		Category:      nil,
//		PriceLessThan: nil,
//	}
//
//	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
//	mockProductRepository.EXPECT().GetProducts(&req).Return([]*domain.Product{
//		&domain.Product{
//			Sku:      nil,
//			Name:     nil,
//			Category: nil,
//			Price: &domain.Price{
//				Original: nil,
//				Final:    nil,
//				Discount: nil,
//				Currency: nil,
//			},
//		},
//	}, nil)
//
//	productService := NewProductService(mockProductRepository)
//	_, err := productService.GetProducts(&req)
//
//	assert.Nil(t, err)
//}

func TestApplyDiscount(t *testing.T) {
	tests := []struct {
		input struct {
			number   int64
			discount int64
		}
		expected int64
	}{
		{
			input: struct {
				number   int64
				discount int64
			}{number: 12050, discount: 30},
			expected: 8435,
		},
	}

	for _, testCase := range tests {
		actual, err := applyDiscount(testCase.input.number, testCase.input.discount)

		assert.Nil(t, err)
		assert.Equal(t, testCase.expected, *actual)
	}
}
