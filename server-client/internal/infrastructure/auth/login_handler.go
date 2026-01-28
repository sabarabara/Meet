package auth

import (
	"log"
	"net/http"
	"server-client/internal/application/dto"
	"server-client/internal/domain/repository/command"
	"server-client/internal/domain/repository/query"
	oidc "server-client/internal/infrastructure/auth/oidc_client"
	strategy "server-client/internal/infrastructure/auth/provider_strategy"
	"time"

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

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	providerName := r.URL.Query().Get("provider")
	if providerName == "" {
		http.Error(w, "Provider not specified", http.StatusBadRequest)
		return
	}

	state := "random-state-string"

	authURL, err := h.providerRegistry.GetAuthURL(providerName, state)
	if err != nil {
		http.Error(w, "Failed to get auth URL", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

func (h *LoginHandler) Callback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := r.URL.Query().Get("code")
	providerName := r.URL.Query().Get("provider")

	token, err := h.providerRegistry.ExchangeAndVerify(ctx, code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusUnauthorized)
		return
	}

	sub := token.Sub

	userid, username, err := h.authQueryRepo.GetAuthenticatedUserInfo(sub, providerName)
	if err != nil {
		log.Printf("failed to get authenticated user info: %v", err)
		return
	}

	//ここはのちに修正する必要あり
	if userid == "" || username == "" {
		dto := dto.NewUserDTO(
			&uuid.Nil,
			"",
			"",
			"",
			"",
			0.0,
		)
		user_dto, err := h.userRepo.UpsertUser(dto)
		if err != nil {
			log.Printf("failed to upsert user: %v", err)
			return
		}
		userid := user_dto.Userid()
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
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		})
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
