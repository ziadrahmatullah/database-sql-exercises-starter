package database

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
