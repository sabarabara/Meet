package query

import (
	"server-client/internal/application/dto"

	"github.com/google/uuid"
)

type UserRepo interface {
	GetUserByID(userID uuid.UUID) (dto.UserDTO, error)
}
