package queries

import (
	"database/sql"
)

func DeleteProduct(db *sql.DB, id int) error {

	sql := "DELETE FROM products WHERE id = $1"

	_, err := db.Query(sql, id)
	if err != nil {
		return err
	}

	return nil
}
