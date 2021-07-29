package app

import (
	"errors"

	"github.com/go-chi/chi/v5"
)

type App struct {
}

func NewApp() (*App, error) {
	return nil, errors.New("not implemented")
}

func (a *App) Router() chi.Router {
	panic("not implemented")
}
