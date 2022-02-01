package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProductsQuery(t *testing.T) {
	builder := GetProductsQuery()

	t.Run("get products without filter", func(t *testing.T) {
		var (
			expectedQuery = fmt.Sprintf("SELECT sku, name, category, price FROM %s", TableProducts)
		)

		query, args, err := builder.ToSql()

		assert.Nil(t, err)
		assert.Nil(t, args)
		assert.Equal(t, expectedQuery, query)
	})

	t.Run("filter by category", func(t *testing.T) {
		var (
			category      = "my_category"
			expectedQuery = fmt.Sprintf("SELECT sku, name, category, price FROM %s "+
				"WHERE category = $1", TableProducts)
			expectedArgs = []interface{}{category}
		)

		filterByCategorySqlBuilder := FilterByCategoryQuery(builder, category)

		query, args, err := filterByCategorySqlBuilder.ToSql()

		assert.Nil(t, err)
		assert.Equal(t, expectedQuery, query)
		assert.Equal(t, expectedArgs, args)
	})

	t.Run("filter by prices less than", func(t *testing.T) {
		var (
			price         int64 = 100
			expectedQuery       = fmt.Sprintf("SELECT sku, name, category, price FROM %s "+
				"WHERE price < $1", TableProducts)
			expectedArgs = []interface{}{price}
		)

		filterByPriceLessThanSqlBuilder := FilterByPriceLessThanQuery(builder, price)

		query, args, err := filterByPriceLessThanSqlBuilder.ToSql()

		assert.Nil(t, err)
		assert.Equal(t, expectedQuery, query)
		assert.Equal(t, expectedArgs, args)
	})

	t.Run("filter by category and price less than", func(t *testing.T) {
		var (
			price         int64 = 100
			category            = "my_category"
			expectedQuery       = fmt.Sprintf("SELECT sku, name, category, price FROM %s "+
				"WHERE category = $1 AND price < $2", TableProducts)
			expectedArgs = []interface{}{category, price}
		)

		filterByCategoryQueryBuilder := FilterByCategoryQuery(builder, category)
		filterByPriceLessThanQueryBuilder := FilterByPriceLessThanQuery(filterByCategoryQueryBuilder, price)

		query, args, err := filterByPriceLessThanQueryBuilder.ToSql()

		assert.Nil(t, err)
		assert.Equal(t, expectedQuery, query)
		assert.Equal(t, expectedArgs, args)
	})

}
