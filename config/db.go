package config

import (
	"database/sql"

	// import sql driver
	_ "github.com/go-sql-driver/mysql"
)

// Connect to database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/resource")
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}
