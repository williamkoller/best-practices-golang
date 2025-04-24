package configs

import (
	"best-practices-golang"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var _, dsn, _ = best_practices_golang.Env()

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
