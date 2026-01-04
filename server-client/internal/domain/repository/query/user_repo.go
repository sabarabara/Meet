package query

import "server-client/internal/application/dto"

type UserRepo interface {
	GetUsers(page int, size int) ([]dto.UserDTO, error)
}
