package strategy

import "context"

type AuthUser struct {
	Sub      string
	Email    string
	Username string
}

type AuthProviderStrategy interface {
	GetName() string
	GetIssuer() string
	Matches(iss string) bool
	GetAuthURL(state string) string
	ExchangeAndVerify(ctx context.Context, code string) (*AuthUser, error)
}
