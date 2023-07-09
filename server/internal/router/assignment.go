package router

import (
	"github.com/fullstackatbrown/auth-infrastructure/internal/handler"
	"github.com/go-chi/chi/v5"
)

func AssignmentRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/", handler.CreateAssignment)
	router.Get("/", handler.GetAssignmentsByEmail)

	router.Route("/{assignmentId}", func(r chi.Router) {
		r.Patch("/", handler.UpdateAssignment)
		r.Delete("/", handler.RemoveAssignment)
	})

	return router
}
