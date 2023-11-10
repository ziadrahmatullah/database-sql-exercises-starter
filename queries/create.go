package queries

import (
	"database-sql-exercises/entity"
	"database/sql"
	"fmt"

	"github.com/shopspring/decimal"
)

func CreateProduct(db *sql.DB, name string, stock int, productCategoryId int64, price decimal.Decimal) (*entity.Product, error) {

	var product entity.Product

	sql := fmt.Sprintf("INSERT INTO products (name, stock, price, category_id, created_at, updated_at) VALUES (%s, %d, %d, %d, NOW(), NOW()) RETURNING id", name, stock, price, productCategoryId)

	err := db.QueryRow(sql).Scan(&product.Id)
	if err != nil {
		return nil, err
	}

	product.Name = name
	product.Stock = stock
	product.Price = price
	product.ProductCategoryId = productCategoryId

	return &product, nil
}
