package auth

import (
	"os"

	"golang.org/x/oauth2/google"
)

type OIDCConfig struct {
	Issuer       string
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Endpoint     interface{}
	Scopes       []string
}

type AuthConfigs struct {
	Google OIDCConfig
}

func LoadAuthConfigs() *AuthConfigs {
	return &AuthConfigs{
		Google: OIDCConfig{
			Issuer:       "https://accounts.google.com",
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
			Endpoint:     google.Endpoint,
			Scopes:       []string{"openid", "email", "profile"},
		},
	}
}
