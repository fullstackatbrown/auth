package auth

import (
	"crypto/sha1"
	"strings"
	"time"

	"github.com/fullstackatbrown/auth-infrastructure/internal/config"
	"github.com/fullstackatbrown/auth-infrastructure/internal/db"
	"github.com/fullstackatbrown/auth-infrastructure/internal/model"
	"github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/avatar"
	"github.com/go-pkgz/auth/provider"
	"github.com/go-pkgz/auth/token"
	"golang.org/x/oauth2"
)

var Service *auth.Service

func defaultOpts() auth.Opts {
	opts := auth.Opts{
		SecretReader: token.SecretFunc(func(id string) (string, error) { // secret key for JWT
			return "secret", nil
		}),
		TokenDuration:   time.Hour * 24 * 14, // token expires in 14 days
		CookieDuration:  time.Hour * 24 * 14, // cookie expires in 14 days
		Issuer:          "fsab-auth",
		DisableXSRF:     true,
		JWTCookieDomain: config.Config.CookieDomain,
		URL:             config.Config.RootUrl,
		AvatarStore:     avatar.NewLocalFS("/tmp"),
		AvatarRoutePath: "/v1/auth/avatars",
		ClaimsUpd: token.ClaimsUpdFunc(func(claims token.Claims) token.Claims { // modify issued token
			if claims.User != nil {
				// check if user is in allowed email domains
				if len(config.Config.AllowedEmailDomains) > 0 {
					for _, domain := range config.Config.AllowedEmailDomains {
						if strings.HasSuffix(claims.User.Email, domain) {
							userLoginHandler(claims.User)
							break
						}
					}
				} else {
					userLoginHandler(claims.User)
				}
			}
			return claims
		}),
	}
	return opts
}

func addGoogleProvider() {
	gClient := auth.Client{
		Cid:     config.Config.OAuth2.ClientID,
		Csecret: config.Config.OAuth2.ClientSecret,
	}

	Service.AddCustomProvider("google", gClient, provider.CustomHandlerOpt{
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/v2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
		InfoURL: "https://www.googleapis.com/oauth2/v3/userinfo",
		MapUserFn: func(data provider.UserData, _ []byte) token.User {
			userInfo := token.User{
				ID:      token.HashID(sha1.New(), data.Value("sub")),
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
}

func init() {
	opts := defaultOpts()
	Service = auth.NewService(opts)

	addGoogleProvider()
}

func userLoginHandler(user *token.User) {
	// TODO attach assignments to user
	dbUser := model.NewUser(user.ID, user.Name, user.Email)
	db.Update(dbUser, true)
}
