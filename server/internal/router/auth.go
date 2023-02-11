package router

import (
	"github.com/go-chi/chi/v5"
)

func AuthRoutes() *chi.Mux {
	router := chi.NewRouter()

	return router
}
