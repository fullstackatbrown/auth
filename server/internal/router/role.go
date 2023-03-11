package router

import (
	"github.com/fullstackatbrown/auth-infrastructure/internal/handler"
	"github.com/go-chi/chi/v5"
)

func RoleRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", handler.ListRoles)

	// TODO: Require admin
	router.Post("/", handler.AddRole)
	router.Delete("/{roleId}", handler.RemoveRole)

	return router
}
