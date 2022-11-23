package server

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/reizt/portfolio/src/adapters/handler"
	"github.com/reizt/portfolio/src/adapters/store"
	"github.com/reizt/portfolio/src/core/app"
	"github.com/reizt/portfolio/src/drivers/db"
)

func setRouter(server *fiber.App) error {
	// Init database
	db, err := db.Connect()
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	// Init store
	store, err := store.Init(db)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	// Init app
	app, err := app.Init(store)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	// Init handler
	h := handler.Init(app)

	// Set routes
	server.Get("/", f(h.SayHello))
	server.Get("/users", f(h.GetUsers))

	return nil
}
