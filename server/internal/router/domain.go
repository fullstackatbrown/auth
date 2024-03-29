package router

import (
	"github.com/fullstackatbrown/auth-infrastructure/internal/handler"
	"github.com/go-chi/chi/v5"
)

func DomainRoutes() *chi.Mux {
	router := chi.NewRouter()

	// TODO: Require super admin
	router.Post("/", handler.CreateDomain)

	router.Route("/{domainName}", func(r chi.Router) {
		r.Get("/", handler.GetDomain)

		// TODO: Require super admin
		r.Delete("/", handler.DeleteDomain)

		r.Mount("/roles", domainRolesRoutes())
	})

	return router
}

func domainRolesRoutes() *chi.Mux {
	router := chi.NewRouter()

	// TODO: Require super admin
	router.Post("/", handler.CreateDomainRole)
	router.Delete("/", handler.DeleteDomainRole)

	return router
}
