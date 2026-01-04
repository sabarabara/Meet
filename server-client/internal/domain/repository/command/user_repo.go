package command

import (
	"server-client/internal/application/dto"

	"github.com/google/uuid"
)

type UserRepo interface {
	UpsertUser(dto dto.UserDTO) (dto.UserDTO, error)
	DeleteUser(userID uuid.UUID) error
}
