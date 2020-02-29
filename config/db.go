package config

import (
	"database/sql"

	// import sql driver
	_ "github.com/go-sql-driver/mysql"
)

// Connect to database
func Connect() error {
	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/experiment")
	if err != nil {
		return err
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}
