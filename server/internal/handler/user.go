package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{"user": "user1"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("User deleted"))
}
