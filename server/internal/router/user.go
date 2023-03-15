package router

import (
	"github.com/fullstackatbrown/auth-infrastructure/internal/handler"
	"github.com/go-chi/chi/v5"
)

func UserRoutes() *chi.Mux {
	// TODO: Require authentication
	router := chi.NewRouter()

	router.Route("/{userId}", func(r chi.Router) {
		r.Get("/", handler.GetUserByEmail)

		// TODO: Require admin
		r.Delete("/", handler.DeleteUser)

		r.Mount("/profile", userProfileRoutes())
		r.Mount("/roles", userRolesRoutes())
	})

	// TODO: Buld upload

	return router
}

func userProfileRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", handler.GetProfile)

	// TODO: require the target user to match the logged in user
	router.Patch("/", handler.UpdateProfile)

	return router
}

func userRolesRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", handler.ListRoles)

	// TODO: Require admin
	router.Post("/", handler.AddRole)
	router.Delete("/", handler.RemoveRole)

	return router
}
