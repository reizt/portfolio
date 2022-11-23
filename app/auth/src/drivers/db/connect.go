package db

import (
	"database/sql"
	"os"
)

func Connect() (*sql.DB, error) {
	SQLITE_PATH := os.Getenv("SQLITE_PATH")
	return sql.Open("sqlite3", SQLITE_PATH)
}
