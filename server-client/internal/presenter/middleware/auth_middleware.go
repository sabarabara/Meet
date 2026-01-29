package middleware

import (
	"net/http"
	"server-client/internal/infrastructure/auth"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	sessionManager *auth.SessionManager
}

func NewAuthMiddleware(sessionManager *auth.SessionManager) *AuthMiddleware {
	return &AuthMiddleware{
		sessionManager: sessionManager,
	}
}

type ctxKey string

const UserKey ctxKey = "user"

func (m *AuthMiddleware) AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("session_id")
		if err != nil {
			prefix := c.GetString("ProxyPrefix")
			c.Redirect(http.StatusTemporaryRedirect, prefix+"/auth/login?provider=google")
			c.Abort()
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
