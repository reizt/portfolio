package server

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func LaunchDevServer(port int) {
	// Init fiber instance
	server := fiber.New()

	// Recover
	server.Use(recoverMw())
	// CORS
	server.Use(corsMw())
	// Logger
	server.Use(loggerMw())

	// Routes
	setRouter(server)

	// Listen
	fmt.Printf("Server started at http://localhost:%d.\n", port)
	log.Fatal(server.Listen(":" + strconv.Itoa(port)))
}
