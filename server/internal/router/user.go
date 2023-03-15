package router

import (
	"github.com/fullstackatbrown/auth-infrastructure/internal/handler"
	"github.com/go-chi/chi/v5"
)

func UserRoutes() *chi.Mux {
	// TODO: Require authentication
	router := chi.NewRouter()

	router.Route("/{userId}", func(r chi.Router) {
		// TODO: Get user by email query param
		r.Get("/", handler.GetUser)

		// TODO: Require admin
		r.Delete("/", handler.DeleteUser)

		r.Mount("/roles", RoleRoutes())
	})

	// TODO: Buld upload

	return router
}
