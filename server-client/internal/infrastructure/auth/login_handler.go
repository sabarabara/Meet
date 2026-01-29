package auth

import (
	"fmt"
	"log"
	"net/http"
	"server-client/internal/application/dto"
	"server-client/internal/domain/repository/command"
	"server-client/internal/domain/repository/query"
	oidc "server-client/internal/infrastructure/auth/oidc_client"
	strategy "server-client/internal/infrastructure/auth/provider_strategy"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoginHandler struct {
	oidcService      *oidc.OIDCService
	sessionManager   *SessionManager
	authQueryRepo    query.AuthRepo
	authCommandRepo  command.AuthRepo
	userRepo         command.UserRepo
	providerRegistry *strategy.ProviderRegistry
}

func NewLoginHandler(
	oidcService *oidc.OIDCService,
	sessionManager *SessionManager,
	authQueryRepo query.AuthRepo,
	authCommandRepo command.AuthRepo,
	userRepo command.UserRepo,
	providerRegistry *strategy.ProviderRegistry,
) *LoginHandler {
	return &LoginHandler{
		oidcService:      oidcService,
		sessionManager:   sessionManager,
		authQueryRepo:    authQueryRepo,
		authCommandRepo:  authCommandRepo,
		userRepo:         userRepo,
		providerRegistry: providerRegistry,
	}
}

func (h *LoginHandler) Login(c *gin.Context) {
	providerName := c.Query("provider")
	if providerName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provider not specified"})
		return
	}

	state := "random-state-string"

	scheme := c.GetString("ProxyScheme")
	prefix := c.GetString("ProxyPrefix")
	host := c.Request.Host

	redirectURI := fmt.Sprintf("%s://%s%s/auth/callback", scheme, host, prefix)
	authURL, err := h.providerRegistry.GetAuthURL(providerName, state, redirectURI)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get auth URL"})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

func (h *LoginHandler) Callback(c *gin.Context) {
	ctx := c.Request.Context()
	code := c.Query("code")
	providerName := c.Query("provider")

	scheme := c.GetString("ProxyScheme")
	prefix := c.GetString("ProxyPrefix")
	host := c.Request.Host
	redirectURI := fmt.Sprintf("%s://%s%s/auth/callback", scheme, host, prefix)

	token, err := h.providerRegistry.ExchangeAndVerify(ctx, code, redirectURI)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to exchange token"})
		return
	}

	sub := token.Sub

	userid, username, err := h.authQueryRepo.GetAuthenticatedUserInfo(sub, providerName)
	if err != nil {
		log.Printf("failed to get authenticated user info: %v", err)
		return
	}

	if userid == "" || username == "" {
		dto := dto.NewUserDTO(
			nil,
			token.Username,
			"",
			"最近始めたばかりのユーザーです。",
			"よろしくね！",
			0.0,
		)
		user_dto, err := h.userRepo.UpsertUser(dto)
		if err != nil {
			log.Printf("failed to upsert user: %v", err)
			return
		}
		userid := user_dto.Userid()
		username := token.Username
		err = h.authCommandRepo.CreateAuthenticatedTable(*userid, sub, providerName)
		if err != nil {
			log.Printf("failed to create authenticated user: %v", err)
			return
		}

		sessionID := uuid.New().String()

		sessData := UserSession{
			UserID:   *userid,
			Username: username,
		}

		if err := h.sessionManager.Create(ctx, sessionID, sessData, 24*time.Hour); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.SetCookie("session_id", sessionID, 3600*24, "/", "", true, true)
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		})
		target := prefix + "/"
		c.Redirect(http.StatusFound, target)
	}
}
