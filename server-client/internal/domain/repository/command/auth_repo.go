package command

import "github.com/google/uuid"

type AuthRepo interface {
	CreateAuthenticatedTable(userid uuid.UUID, sub string, provider string) error
}
