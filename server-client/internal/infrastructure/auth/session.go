package auth

import (
	"context"
	redis "server-client/pkg/db"
	"time"

	"github.com/google/uuid"
)

type SessionManager struct {
	client *redis.RedisClient
}

type UserSession struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
}

func NewSessionManager(client *redis.RedisClient) *SessionManager {
	return &SessionManager{client: client}
}

func (s *SessionManager) Create(ctx context.Context, sessionID string, user UserSession, duration time.Duration) error {
	return s.client.Set(ctx, sessionID, user, duration)
}

func (s *SessionManager) Get(ctx context.Context, sessionID string) (*UserSession, error) {

	val, err := s.client.Get(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, nil
	}

	return &UserSession{
		UserID:   uuid.MustParse(val["user_id"]),
		Username: val["username"],
	}, nil
}

func (s *SessionManager) Delete(ctx context.Context, sessionID string) error {
	return s.client.Delete(ctx, sessionID)
}
