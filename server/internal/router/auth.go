package router

import (
	"crypto/sha1"
	"net/http"
	"strings"
	"time"

	"github.com/fullstackatbrown/auth-infrastructure/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/avatar"
	"github.com/go-pkgz/auth/provider"
	"github.com/go-pkgz/auth/token"
	"golang.org/x/oauth2"
)

func DefaultAuthService() *auth.Service {
	opts := auth.Opts{
		SecretReader: token.SecretFunc(func(id string) (string, error) { // secret key for JWT
			return "secret", nil
		}),
		TokenDuration:  time.Hour * 24 * 14, // token expires in 14 days
		CookieDuration: time.Hour * 24 * 14, // cookie expires in 14 days
		Issuer:         "fsab-auth",
		DisableXSRF:    true,
		// JWTCookieName:  "fsab-auth",
		SecureCookies:  true,
		SameSiteCookie: http.SameSiteNoneMode,
		URL:            config.Config.RootUrl,
		AvatarStore:    avatar.NewLocalFS("/tmp"),
		ClaimsUpd: token.ClaimsUpdFunc(func(claims token.Claims) token.Claims { // modify issued token
			if claims.User != nil {
				// check if user is in allowed email domains
				if len(config.Config.AllowedEmailDomains) > 0 {
					for _, domain := range config.Config.AllowedEmailDomains {
						if strings.HasSuffix(claims.User.Email, domain) {
							// TODO save to db and add roles and stuff
							return claims
						}
					}
					return claims // don't save to db
				}
			}
			return claims // don't save to db
		}),
	}

	service := auth.NewService(opts)

	gClient := auth.Client{
		Cid:     config.Config.OAuth2.ClientID,
		Csecret: config.Config.OAuth2.ClientSecret,
	}

	service.AddCustomProvider("google", gClient, provider.CustomHandlerOpt{
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/v2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
		InfoURL: "https://www.googleapis.com/oauth2/v3/userinfo",
		MapUserFn: func(data provider.UserData, _ []byte) token.User {
			userInfo := token.User{
				ID: "google_" + token.HashID(sha1.New(),
					data.Value("sub")),
				Name:    data.Value("name"),
				Email:   data.Value("email"),
				Picture: data.Value("picture"),
			}
			// fail if email is not in AllowedEmailDomains
			if len(config.Config.AllowedEmailDomains) > 0 {
				for _, domain := range config.Config.AllowedEmailDomains {
					if strings.HasSuffix(userInfo.Email, domain) {
						return userInfo
					}
				}
				return token.User{}
			}
			return userInfo
		},
		Scopes: []string{"profile", "email"},
	})

	return service
}

func AuthRoutes() *chi.Mux {
	router := chi.NewRouter()

	service := DefaultAuthService()
	authRoutes, avaRoutes := service.Handlers()

	router.Mount("/", authRoutes)
	router.Mount("/avatar", avaRoutes)

	return router
}
