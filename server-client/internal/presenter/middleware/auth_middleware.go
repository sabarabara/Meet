package middleware

import (
	"net/http"
	"server-client/internal/infrastructure/auth"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	sessionManager *auth.SessionManager
}

type ctxKey string

const UserKey ctxKey = "user"

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("session_id")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		userSession, err := m.sessionManager.Get(c.Request.Context(), cookie)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session expired"})
			return
		}
		c.Set("user", userSession)

		c.Next()
	}
}
