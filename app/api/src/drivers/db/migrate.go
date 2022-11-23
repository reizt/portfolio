package db

import (
	"log"

	"github.com/reizt/portfolio/src/adapters/store"
)

func MigrateSync() error {
	db, err := Connect()
	if err != nil {
		log.Fatal("failed to connect database")
		log.Fatal(err.Error())
		return err
	}

	if err := store.MigrateSync(db); err != nil {
		log.Fatal("failed to exec migration")
		log.Fatal(err.Error())
		return err
	}
	return nil
}
