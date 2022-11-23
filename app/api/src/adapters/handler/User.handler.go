package handler

import (
	"github.com/reizt/portfolio/src/adapters/handler/spec"
	"github.com/reizt/portfolio/src/core"
	"github.com/reizt/portfolio/src/core/app"
)

func (h *handler) GetUsers(c Ctx) error {
	// Make input
	input := app.GetUsersInput{}

	// Do action
	users, err := h.app.GetUsers(input)

	// Make output
	if err != nil {
		if err == core.ErrInvalidInput {
			return (*c.Status(400)).SendJson(map[string]string{"sentence": "invalid input"})
		}
		return (*c.Status(500)).SendJson(map[string]string{"sentence": "An error occured!"})
	}
	data := spec.GetUsers200Response(*users)
	return c.SendJson(data)
}
