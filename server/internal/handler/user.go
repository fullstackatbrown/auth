package handler

import (
	"net/http"

	"github.com/fullstackatbrown/auth-infrastructure/internal/db"
	"github.com/go-chi/render"
)

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
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
}

func AddUserRole(w http.ResponseWriter, r *http.Request) {
}

func RemoveUserRole(w http.ResponseWriter, r *http.Request) {
}
