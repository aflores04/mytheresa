package db

import (
	"github.com/Masterminds/squirrel"
)

const (
	TableProducts = "products"
)

// will generate dollar statement in queries like $1, $2 ... instead of ? (deprecated)
var (
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

func GetProductsQuery() squirrel.SelectBuilder {
	builder := psql.
		Select("sku, name, category, price").
		From(TableProducts)

	return builder
}

func FilterByCategoryQuery(builder squirrel.SelectBuilder, category string) squirrel.SelectBuilder {
	builder = builder.Where(squirrel.Eq{"category": category})

	return builder
}

func FilterByPriceLessThanQuery(builder squirrel.SelectBuilder, price int64) squirrel.SelectBuilder {
	builder = builder.Where(squirrel.Lt{"price": price})

	return builder
}
