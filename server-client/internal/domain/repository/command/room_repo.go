package command

import (
	"server-client/internal/application/dto"

	"github.com/google/uuid"
)

type RoomRepo interface {
	InsertRoom(dto dto.RoomDTO) error
	DeleteRoom(roomid uuid.UUID) error
}
