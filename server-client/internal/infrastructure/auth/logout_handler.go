package auth

import (
	"log"
	"net/http"
)

type LogoutHandler struct {
	sessionManager *SessionManager
}

func NewLogoutHandler(sm *SessionManager) *LogoutHandler {
	return &LogoutHandler{
		sessionManager: sm,
	}
}

func (h *LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	sessionID := cookie.Value

	if err := h.sessionManager.Delete(r.Context(), sessionID); err != nil {
		log.Printf("failed to delete session from redis: %v", err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, "/login", http.StatusFound)
}
