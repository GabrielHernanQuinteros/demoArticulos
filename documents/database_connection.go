package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectarDB() (*sql.DB, error) {
	return sql.Open("mysql", ConnectionString)
}
