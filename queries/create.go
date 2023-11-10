package queries

import (
	"database-sql-exercises/entity"
	"database/sql"
	"log"

	"github.com/shopspring/decimal"
)

func CreateProduct(db *sql.DB, name string, stock int, productCategoryId int64, price decimal.Decimal) (*entity.Product, error) {

	var product entity.Product

	stmt, err := db.Prepare("INSERT INTO products (category_id, name, stock, price, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(productCategoryId, name, stock, price)
	if err != nil {
		return nil, err
	}

	product.Name = name
	product.Stock = stock
	product.Price = price
	product.ProductCategoryId = productCategoryId

	return &product, nil
}
