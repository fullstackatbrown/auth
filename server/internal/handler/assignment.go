package handler

import (
	"net/http"

	"github.com/fullstackatbrown/auth-infrastructure/internal/db"
	"github.com/fullstackatbrown/auth-infrastructure/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	// decode the request body into a new Assignment
	newAssignment := &model.Assignment{}
	err := render.DecodeJSON(r.Body, newAssignment)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"message": "invalid request body"})
		return
	}

	// insert the new Assignment into the db
	err = db.Update(newAssignment, true)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	render.Status(r, http.StatusCreated)
}

func GetAssignmentsByEmail(w http.ResponseWriter, r *http.Request) {
	// get the email from query params
	email := r.URL.Query().Get("email")

	// find the Assignments in the db
	assignments, err := db.FindAssignmentsByEmail(email)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	// return the Assignments
	render.Status(r, http.StatusOK)
	render.JSON(w, r, assignments)
}

func RemoveAssignment(w http.ResponseWriter, r *http.Request) {
	// get the assignmentId from path params
	assignmentId := chi.URLParam(r, "assignmentId")

	// find the Assignment in the db
	assignment, err := db.FindAssignmentById(assignmentId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"message": "assignment not found"})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	// delete the Assignment from the db
	err = db.Delete(assignment)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	render.Status(r, http.StatusNoContent)
}

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	// get the assignmentId from path params
	assignmentId := chi.URLParam(r, "assignmentId")

	// find the Assignment in the db
	assignment, err := db.FindAssignmentById(assignmentId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"message": "assignment not found"})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	// decode the request body into the found assignment
	err = render.DecodeJSON(r.Body, &assignment)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"message": "invalid request body"})
		return
	}

	// update the Assignment in the db
	err = db.Update(assignment, false)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	render.Status(r, http.StatusOK)
}
