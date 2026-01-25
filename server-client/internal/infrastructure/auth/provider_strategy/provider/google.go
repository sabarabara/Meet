package provider

import (
	"context"
	oidc "server-client/internal/infrastructure/auth/oidc_client"
	auth "server-client/internal/infrastructure/auth/provider_strategy/interface"
)

type GoogleStrategy struct {
	OIDCService *oidc.OIDCService
}

func NewGoogleStrategy(svc *oidc.OIDCService) *GoogleStrategy {
	return &GoogleStrategy{OIDCService: svc}
}
func (s *GoogleStrategy) GetName() string   { return "google" }
func (s *GoogleStrategy) GetIssuer() string { return "https://accounts.google.com" }
func (s *GoogleStrategy) Matches(iss string) bool {
	return iss == "https://accounts.google.com" || iss == "accounts.google.com"
}

func (s *GoogleStrategy) GetAuthURL(state string) string {
	return s.OIDCService.GetAuthURL(state)
}

func (s *GoogleStrategy) ExchangeAndVerify(ctx context.Context, code string) (*auth.AuthUser, error) {
	idToken, err := s.OIDCService.ExchangeAndVerify(ctx, code)
	if err != nil {
		return nil, err
	}

	var claims struct {
		Sub   string `json:"sub"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	if err := idToken.Claims(&claims); err != nil {
		return nil, err
	}

	return &auth.AuthUser{
		Sub:      claims.Sub,
		Email:    claims.Email,
		Username: claims.Name,
	}, nil
}
