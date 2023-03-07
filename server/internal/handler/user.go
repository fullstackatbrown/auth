package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{"user": "user1"})
}

// TODO: Add preassigned roles
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	w.Write([]byte("User created"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("User deleted"))
}
