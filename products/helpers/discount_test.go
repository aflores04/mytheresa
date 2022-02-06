package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
		{
			input: struct {
				number   int64
				discount int64
			}{number: 89000, discount: 30},
			expected: 62300,
		},
	}

	for _, testCase := range tests {
		actual := ApplyDiscount(testCase.input.number, testCase.input.discount)

		assert.Equal(t, testCase.expected, actual)
	}
}
