package config

import (
	"database/sql"
	"fmt"

	// import sql driver
	_ "github.com/go-sql-driver/mysql"
)

// Connect to database
func Connect(cfg Configuration) (*sql.DB, error) {
	s := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s",
		cfg.Database.DBUser,
		cfg.Database.DBPass,
		cfg.Database.DBProtocol,
		cfg.Database.DBHost,
		cfg.Database.DBPort,
		cfg.Database.DBName,
	)
	db, err := sql.Open("mysql", s)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}
