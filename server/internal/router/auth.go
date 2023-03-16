package router

import (
	a "github.com/fullstackatbrown/auth-infrastructure/internal/auth"
	"github.com/go-chi/chi/v5"
)

func AuthRoutes() *chi.Mux {
	router := chi.NewRouter()

	auth, ava := a.Service.Handlers()

	router.Mount("/", auth)
	router.Mount("/avatars", ava)

	return router
}
