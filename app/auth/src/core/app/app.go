package app

import (
	"errors"

	"github.com/reizt/portfolio/src/core/istore"
)

type App struct {
	Store istore.Store
}

func Init(store *istore.Store) (*App, error) {
	if store == nil {
		return nil, errors.New("error")
	}
	return &App{*store}, nil
}
