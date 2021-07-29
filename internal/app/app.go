package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/qbarrand/arbalete/internal/models"
	"gorm.io/gorm"
)

type App struct {
	db *gorm.DB
}

func NewApp(db *gorm.DB) (*App, error) {
	if err := db.AutoMigrate(models.AllModels()...); err != nil {
		return nil, err
	}

	return &App{db: db}, nil
}

func (a *App) Router() chi.Router {
	r := chi.NewRouter()

	r.Route("/trains", func(r chi.Router) {
		r.Method(http.MethodGet, "/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte("some trains"))
		}))
	})

	return r
}
