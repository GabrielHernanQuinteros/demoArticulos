package routes

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func fnConectarDB() (*sql.DB, error) {
	return sql.Open("mysql", ConnectionString)
}
