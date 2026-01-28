package middleware

import (
	"context"
	"net/http"
	"server-client/internal/infrastructure/auth"
)

type AuthMiddleware struct {
	sessionManager *auth.SessionManager
}

type ctxKey string

const UserKey ctxKey = "user"

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userSession, err := m.sessionManager.Get(r.Context(), cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized: Session expired", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, userSession)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
