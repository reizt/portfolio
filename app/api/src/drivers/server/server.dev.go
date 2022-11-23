package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func LaunchDevServer() {
	// Init fiber instance
	server := fiber.New()

	// CORS
	server.Use(corsMw())
	// Logger
	server.Use(loggerMw())

	// Routes
	setRouter(server)

	// Listen
	port := getDevPort()
	fmt.Printf("Server started at http://localhost:%s.\n", port)
	log.Fatal(server.Listen(":" + port))
}

func getDevPort() string {
	port := os.Getenv("APP_LISTEN_PORT")
	if port == "" {
		panic("Port is not specified as an environment variable $APP_LISTEN_PORT.")
	}
	return port
}
