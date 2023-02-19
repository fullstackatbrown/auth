package handler

import (
	"fmt"
	"net/http"
)

func PublicRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Public Route")
}

func ProtectedRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Protected Route")
}
