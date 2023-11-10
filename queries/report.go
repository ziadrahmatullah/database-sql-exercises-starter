package queries

import (
	"database-sql-exercises/entity"
	"database/sql"
)

func queryDB(db *sql.DB, query string, report *int) (error){
	rows, err := db.Query(query)
	if err != nil {
		return  err
	}
	for rows.Next(){
		err = rows.Scan(report)
		if err != nil {
			return  err
		}
	}

	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func ProductReport(db *sql.DB) (*entity.Report, error) {

	expensiveProduct := `SELECT count(id) as product from products where price >= 10000`
	goodDealProduct := `Select count(id) as product from products where price between 5001 and 9999`
	cheapProduct := `select count(id) as product from products where price <= 5000`

	var report entity.Report
	err := queryDB(db, expensiveProduct, &report.ExpensiveProduct)
	if err != nil {
		return nil, err
	}

	err = queryDB(db, goodDealProduct, &report.GoodDealProduct)
	if err != nil {
		return nil, err
	}

	err = queryDB(db, cheapProduct, &report.CheapProduct)
	if err != nil {
		return nil, err
	}

	return &report, nil
}
