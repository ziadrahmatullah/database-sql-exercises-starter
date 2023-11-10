package entity

import "github.com/shopspring/decimal"

type Product struct {
	Id                int64
	Name              string
	Price             decimal.Decimal
	ProductCategoryId int64
	Stock             int
}
