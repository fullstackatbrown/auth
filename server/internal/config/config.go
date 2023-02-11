package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var Config *ServerConfig

// ServerConfig is a struct that contains configuration values for the server.
type ServerConfig struct {
	// AllowedOrigins is a list of URLs that the server will accept requests from.
	AllowedOrigins []string
	// AllowedEmailDomains is a list of email domains that the server will allow account registrations from. If empty,
	// all domains will be allowed.
	AllowedEmailDomains []string
	// IsHTTPS should be set to true for production.
	IsHTTPS bool
	// SessionCookieName is the name to use for the session cookie.
	SessionCookieName string
	// SessionCookieExpiration is the amount of time a session cookie is valid. Max 5 days.
	SessionCookieExpiration time.Duration
	// Port is the port the server should run on.
	Port int
	// Google OAuth2 config
	OAuth2 *oauth2.Config
}

func DefaultDevelopmentConfig() *ServerConfig {
	env, err := godotenv.Read()
	if err != nil {
		log.Panic("Error loading env file.")
	}

	oauth := &oauth2.Config{
		RedirectURL:  "http://localhost:8000/callback",
		ClientID:     env["GOOGLE_CLIENT_ID"],
		ClientSecret: env["GOOGLE_CLIENT_SECRET"],
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	return &ServerConfig{
		AllowedOrigins:          []string{"http://localhost:3000"},
		AllowedEmailDomains:     []string{"brown.edu", "gmail.com"},
		IsHTTPS:                 false,
		SessionCookieName:       "fsab-session",
		SessionCookieExpiration: time.Hour * 24 * 14,
		Port:                    8000,
		OAuth2:                  oauth,
	}
}

func init() {
	Config = DefaultDevelopmentConfig()
}
