package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// CORS
	AllowOrigins := os.Getenv("CORS_ALLOW_ORIGINS")
	app.Use(cors.New(cors.Config{
		AllowOrigins: AllowOrigins,
		AllowMethods: "GET,POST,PATCH,PUT,DELETE",
	}))

	// Logger
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world from portfolio API!")
	})

	port := os.Getenv("APP_LISTEN_PORT")
	if port == "" {
		log.Fatal("Port is not specified as an environment variable $APP_LISTEN_PORT.")
		return
	}
	fmt.Printf("Server started at http://localhost:%s.\n", port)
	log.Fatal(app.Listen(":" + port))
}
