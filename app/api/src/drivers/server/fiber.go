package server

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"

	"github.com/reizt/portfolio/src/adapters/handler"
	"github.com/reizt/portfolio/src/utils"
)

// f converts returns a handler that converts fiber.Ctx to handler.Ctx and exec handler with it.
func f(handler func(c handler.Ctx) error) fiber.Handler {
	return func(c *fiber.Ctx) error { return handler(fiberImpl{c}) }
}

type fiberImpl struct {
	fiberCtx *fiber.Ctx
}

func (i fiberImpl) Body() (*([]byte), error) {
	bytes := i.fiberCtx.Body()
	return &bytes, nil
}

func (i fiberImpl) Status(status int) *handler.Ctx {
	i.fiberCtx.Status(status)
	return utils.Ptr(handler.Ctx(i))
}

func (i fiberImpl) SendJson(body any) error {
	bytes, err := json.Marshal(body)
	if err != nil {
		i.fiberCtx.Status(500)
		return i.fiberCtx.SendString("An error occured!")
	}
	return i.fiberCtx.Send(bytes)
}
