package config

import (
	"os"
	"strings"
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
	// CookieName is the name to use for the session cookie.
	CookieName string
	// CookieExpiration is the amount of time a session cookie is valid. Max 5 days.
	CookieExpiration time.Duration
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
	// Cookie Domain
	CookieDomain string
}

func init() {
	godotenv.Load()

	oauth := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	Config = &ServerConfig{
		AllowedOrigins:      strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),
		AllowedEmailDomains: strings.Split(os.Getenv("ALLOWED_EMAIL_DOMAINS"), ","),
		IsHTTPS:             os.Getenv("IS_HTTPS") == "true",
		CookieName:          os.Getenv("COOKIE_NAME"),
		CookieExpiration:    time.Hour * 24 * 14,
		Port:                os.Getenv("SERVER_PORT"),
		OAuth2:              oauth,
		MongoUri:            os.Getenv("MONGO_URI"),
		DbName:              os.Getenv("DB_NAME"),
		RootUrl:             os.Getenv("AUTH_ROOT_URL"),
		CookieDomain:        os.Getenv("COOKIE_DOMAIN"),
	}
}
