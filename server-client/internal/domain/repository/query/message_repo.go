package query

import (
	"server-client/internal/application/dto"

	"github.com/google/uuid"
)

type MessageRepo interface {
	GetMessages(roomid uuid.UUID, page int, size int) ([]dto.MessageDTO, error)
}
