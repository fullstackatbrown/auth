package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fullstackatbrown/auth-infrastructure/internal/config"
	"github.com/fullstackatbrown/auth-infrastructure/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/avatar"
	"github.com/go-pkgz/auth/token"
	"github.com/rs/cors"
)

func Start() {
	if config.Config == nil {
		log.Panic("Missing configuration!")
	}

	// define options
	authOptions := auth.Opts{
		SecretReader: token.SecretFunc(func(id string) (string, error) { // secret key for JWT
			return "secret", nil
		}),
		TokenDuration:  time.Hour * 24 * 14, // token expires in 14 days
		CookieDuration: time.Hour * 24 * 14, // cookie expires in 14 days
		Issuer:         "fsab-auth",
		URL:            fmt.Sprintf("http://localhost:%v", config.Config.Port),
		AvatarStore:    avatar.NewLocalFS("/tmp"),
		// Validator: token.ValidatorFunc(func(_ string, claims token.Claims) bool {
		// 	// allow only @brown.edu emails
		// 	if len(config.Config.AllowedEmailDomains) > 0 {
		// 		for _, domain := range config.Config.AllowedEmailDomains {
		// 			if strings.HasSuffix(claims.Email, domain) {
		// 				return true
		// 			}
		// 		}
		// 		return false
		// 	}
		// 	return true
		// }),
	}

	// create auth service with providers
	authService := auth.NewService(authOptions)
	authService.AddProvider("google", config.Config.OAuth2.ClientID, config.Config.OAuth2.ClientSecret)

	m := authService.Middleware()

	// setup http server
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
	)
	router.Get("/open", handler.PublicRoute)                    // open api
	router.With(m.Auth).Get("/private", handler.ProtectedRoute) // protected api

	// setup auth routes
	authRoutes, avaRoutes := authService.Handlers()
	router.Mount("/auth", authRoutes)  // add auth handlers
	router.Mount("/avatar", avaRoutes) // add avatar handler

	c := cors.New(cors.Options{
		AllowedOrigins:   config.Config.AllowedOrigins,
		AllowedHeaders:   []string{"Cookie", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH"},
		ExposedHeaders:   []string{"Set-Cookie"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	log.Printf("Server is listening on port %v\n", config.Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Config.Port), handler))
}
