package store

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/reizt/portfolio/src/core/istore"
)

var (
	userTableName   = "User"
	userColumnNames = struct {
		ID    string
		Email string
	}{
		ID:    "ID",
		Email: "Email",
	}
)

func Init(db *sql.DB) (*istore.Store, error) {
	user := userStore{db}

	store := istore.Store{
		User: user,
	}
	return &store, nil
}

func MigrateSync(db *sql.DB) error {
	// Make query
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			%s VARCHAR(255) NOT NULL,
			%s VARCHAR(255) NOT NULL,
			PRIMARY KEY (%s)
		)
	`, userTableName, userColumnNames.ID, userColumnNames.Email, userColumnNames.ID)

	// Check query
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	// Exec query
	if _, err := stmt.Exec(); err != nil {
		return err
	}

	return nil
}
