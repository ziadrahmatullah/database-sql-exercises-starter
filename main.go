package main

import (
	"database-sql-exercises/database"
	"database-sql-exercises/queries"
	"fmt"
	"log"

	"github.com/shopspring/decimal"
)

func main() {
	// connecting to DB
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("unable to connect to the database: %v", err)
	}

	products, err := queries.ListProducts(db)
	if err != nil {
		log.Fatalf("unable to list products: %v", err)
	}

	// list products
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
	fmt.Println("Successfully created a new product:")
	fmt.Println("===================================")
	fmt.Printf("%s, Rp.%s, %d unit(s)\n", product.Name, product.Price.String(), product.Stock)

	// delete a product
	err = queries.DeleteProduct(db, 1)
	if err != nil {
		log.Fatalf("unable to delete a product: %v", err)
	}
	fmt.Println("Successfully deleted a product:")
	fmt.Println("===================================")

	// implement your report product here...
}
