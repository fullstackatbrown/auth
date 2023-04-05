package handler

import (
	"net/http"

	"github.com/fullstackatbrown/auth-infrastructure/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsersByEmail(w http.ResponseWriter, r *http.Request) {
	// get email from query param
	email := r.URL.Query().Get("email")

	// handle error
	users, _ := db.FindUsersByEmail(email)

	render.JSON(w, r, users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
}

func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
}

func ListUserRoles(w http.ResponseWriter, r *http.Request) {
	// get user id from path param
	userId := chi.URLParam(r, "userId")

	// get user object from db
	user, err := db.FindUserById(userId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"message": "user not found"})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "internal server error"})
		return
	}

	// get roles from user object
	roles := user.Roles

	render.JSON(w, r, roles)
}

func AddUserRole(w http.ResponseWriter, r *http.Request) {
}

func RemoveUserRole(w http.ResponseWriter, r *http.Request) {
}
