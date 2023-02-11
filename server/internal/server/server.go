package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fullstackatbrown/auth-infrastructure/internal/config"
	"github.com/fullstackatbrown/auth-infrastructure/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
	)

	router.HandleFunc("/login", handler.GoogleLogin)
	router.HandleFunc("/callback", handler.GoogleCallback)

	return router
}

func Start() {
	if config.Config == nil {
		log.Panic("Missing configuration!")
	}

	router := Routes()
	c := cors.New(cors.Options{
		// AllowedOrigins:   config.Config.AllowedOrigins,
		AllowedHeaders:   []string{"Cookie", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH"},
		ExposedHeaders:   []string{"Set-Cookie"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	log.Printf("Server is listening on port %v\n", config.Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Config.Port), handler))
}
