package queries

import (
	"database-sql-exercises/entity"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func BulkInsert(db *sql.DB, products []*entity.Product) error {
    // valueStrings := make([]string, 0, len(products))
    // valueArgs := make([]interface{}, 0, len(products) * 3)
    // i := 0
    // for _, product := range products {
    //     valueStrings = append(valueStrings, fmt.Sprintf("(@p%d, @p%d, @p%d, @p%d)", i*3+1, i*3+2, i*3+3, i*3+4))
    //     valueArgs = append(valueArgs, product.Name)
    //     valueArgs = append(valueArgs, product.Price)
    //     valueArgs = append(valueArgs, product.ProductCategoryId)
    //     valueArgs = append(valueArgs, product.Stock)
    //     i++
    // }
    // sqlQuery := fmt.Sprintf("INSERT INTO products (name, price, category_id, stock) VALUES %s", strings.Join(valueStrings, ","))

    // var params []interface{}

    // for i := 0; i < len(valueArgs); i++ {
    //     var param sql.NamedArg
    //     param.Name = fmt.Sprintf("p%v", i+1)
    //     param.Value = valueArgs[i]
    //     params = append(params, param)
    // }

    // _, err := db.Exec(sqlQuery, params...)
    // return err
	valueStrings := make([]string, 0, len(products))
    valueArgs := make([]interface{}, 0, len(products)*4) // Updated to *4 for each parameter
    i := 0
    for _, product := range products {
        valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)", i*6+1, i*6+2, i*6+3, i*6+4, i*6+5, i*6+6))
        valueArgs = append(valueArgs, product.Name)
        valueArgs = append(valueArgs, product.Price)
        valueArgs = append(valueArgs, product.ProductCategoryId)
        valueArgs = append(valueArgs, product.Stock)
        valueArgs = append(valueArgs, time.Now())
        valueArgs = append(valueArgs, time.Now())
        i++
    }
    sqlQuery := fmt.Sprintf("INSERT INTO products (name, price, category_id, stock, created_at, updated_at) VALUES %s", strings.Join(valueStrings, ","))

    _, err := db.Exec(sqlQuery, valueArgs...)
    return err
}