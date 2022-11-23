package handler

import (
	"github.com/reizt/portfolio/src/core/app"
)

type handler struct {
	app *app.App
}

func Init(app *app.App) *handler {
	return &handler{app}
}
