package router

import (
	"github.com/fullstackatbrown/auth-infrastructure/internal/handler"
	"github.com/go-chi/chi/v5"
)

func AssignmentRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/", handler.CreateAssignment)
	router.Get("/", handler.ListAssignments)
	router.Route("/{assignmentId}", func(r chi.Router) {
		r.Get("/", handler.GetAssignment)

		r.Delete("/", handler.RemoveAssignment)
		r.Patch("/", handler.UpdateAssignment)
	})

	return router
}
