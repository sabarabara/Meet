package strategy

import (
	"context"
	"errors"
	"fmt"
	strategy "server-client/internal/infrastructure/auth/provider_strategy/interface"
	provider "server-client/internal/infrastructure/auth/provider_strategy/provider"

	oidc "server-client/internal/infrastructure/auth/oidc_client"
)

type ProviderRegistry struct {
	strategies []strategy.AuthProviderStrategy
}

func NewProviderRegistry(svc *oidc.OIDCService) *ProviderRegistry {
	return &ProviderRegistry{
		strategies: []strategy.AuthProviderStrategy{
			&provider.GoogleStrategy{
				OIDCService: svc,
			},
			// ここに新しく追加したいプロバイダを追加してね!!
			//ex : &FacebookStrategy{},
		},
	}
}

func (r *ProviderRegistry) Get(name string) (strategy.AuthProviderStrategy, error) {
	for _, s := range r.strategies {
		if s.GetName() == name {
			return s, nil
		}
	}
	return nil, errors.New("provider not found")
}

func (r *ProviderRegistry) ExchangeAndVerify(ctx context.Context, code string) (*strategy.AuthUser, error) {
	for _, s := range r.strategies {
		user, err := s.ExchangeAndVerify(ctx, code)
		if err == nil {
			return user, nil
		}
	}
	return nil, errors.New("failed to exchange and verify token with all providers")
}

func (r *ProviderRegistry) GetAuthURL(name string, state string) (string, error) {
	for _, s := range r.strategies {
		if s.GetName() == name {
			fmt.Println(name)
			return s.GetAuthURL(state), nil
		}
	}
	return "", errors.New("provider not found")
}
