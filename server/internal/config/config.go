package config

import (
	"os"
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
	Port string
	// Google OAuth2 config
	OAuth2 *oauth2.Config
	// MongoDB Atlas cluster URI
	MongoUri string
	// Database name
	DbName string
	// Root URL
	RootUrl string
}

func DefaultDevelopmentConfig() *ServerConfig {
	godotenv.Load()

	oauth := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	return &ServerConfig{
		AllowedOrigins:          []string{"localhost:3000", "localhost:8080", "https://here-backend.up.railway.app"},
		AllowedEmailDomains:     []string{"@brown.edu"},
		IsHTTPS:                 false,
		SessionCookieName:       "fsab-session",
		SessionCookieExpiration: time.Hour * 24 * 14,
		Port:                    os.Getenv("PORT"),
		OAuth2:                  oauth,
		MongoUri:                os.Getenv("MONGO_URI"),
		DbName:                  os.Getenv("DB_NAME"),
		RootUrl:                 os.Getenv("ROOT_URL"),
	}
}

func init() {
	Config = DefaultDevelopmentConfig()
}
