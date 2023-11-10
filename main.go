package main

import (
	"database-sql-exercises/database"
	"database-sql-exercises/entity"
	"database-sql-exercises/queries"
	"fmt"
	"log"

	"github.com/shopspring/decimal"
)

// DATABASE_URL="postgres://postgres:postgres@localhost:5432/shop_db" go run .
func main() {
	// connecting to DB
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("unable to connect to the database: %v", err)
	}
	defer db.Close()

	products, err := queries.ListProducts(db)
	if err != nil {
		log.Fatalf("unable to list products: %v", err)
	}

	// list products
	fmt.Println("=================")
	fmt.Println("List of products:")
	fmt.Println("=================")
	for i := 0; i < len(products); i++ {
		product := products[i]
		fmt.Printf("%d. %s, Rp.%s, %d unit(s)\n", (i + 1), product.Name, product.Price.String(), product.Stock)
	}

	// create a new product
	product, err := queries.CreateProduct(db, "digital clock", 100, 3, decimal.NewFromInt(10000))
	if err != nil {
		log.Fatalf("unable to create a new product: %v", err)
	}
	fmt.Println("===================================")
	fmt.Println("Successfully created a new product:")
	fmt.Println("===================================")
	fmt.Printf("%s, Rp.%s, %d unit(s)\n", product.Name, product.Price.String(), product.Stock)

	// delete a product
	err = queries.DeleteProduct(db, 1)
	if err != nil {
		log.Fatalf("unable to delete a product: %v", err)
	}
	fmt.Println("===================================")
	fmt.Println("Successfully deleted a product:")
	fmt.Println("===================================")

	// implement your report product here...
	report, err := queries.ProductReport(db)
	if err != nil {
		log.Fatalf("unable to create a report: %v", err)
	}
	fmt.Println("===================================")
	fmt.Println("Product Report:")
	fmt.Println("===================================")
	fmt.Printf("1. Expensive products: %d product(s)\n", report.ExpensiveProduct)
	fmt.Printf("2. Good Deal products: %d product(s)\n", report.GoodDealProduct)
	fmt.Printf("3. Cheap products: %d product(s)\n", report.CheapProduct)

	// Bulk insert
	var productsBulk []*entity.Product
	product1 := &entity.Product{Name: "Ubi Manis", Price: decimal.NewFromInt32(10000), ProductCategoryId: 4, Stock: 10}
	product2 := &entity.Product{Name: "Singkong", Price: decimal.NewFromInt32(1000), ProductCategoryId: 4, Stock: 15}
	product3 := &entity.Product{Name: "Kerupuk Udang", Price: decimal.NewFromInt32(5000), ProductCategoryId: 4, Stock: 5}
	productsBulk = append(productsBulk, product1)
	productsBulk = append(productsBulk, product2)
	productsBulk = append(productsBulk, product3)
	err = queries.BulkInsert(db, productsBulk)
	if err != nil {
		log.Fatalf("unable to create a new product: %v", err)
	}
	fmt.Println("===================================")
	fmt.Println("Successfully created a new product:")
	fmt.Println("===================================")
	for _ ,val := range productsBulk{
		fmt.Printf("Product Category: %d, Name: %s, Stock: %d, Price: %d\n", val.ProductCategoryId, val.Name, val.Stock, val.Price.IntPart())
	}	
}
