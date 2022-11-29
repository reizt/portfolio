package server

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func corsMw() func(c *fiber.Ctx) error {
	AllowOrigins := os.Getenv("CORS_ALLOW_ORIGINS")
	AllowMethods := "GET,POST,PATCH,PUT,DELETE"

	return cors.New(cors.Config{
		AllowOrigins: AllowOrigins,
		AllowMethods: AllowMethods,
	})
}

func loggerMw() func(c *fiber.Ctx) error {
	return logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	})
}

func recoverMw() func(c *fiber.Ctx) error {
	return recover.New()
}
