package query

import "server-client/internal/application/dto"

type RoomRepo interface {
	GetRooms(page int, size int) ([]dto.RoomDTO, error)
}
