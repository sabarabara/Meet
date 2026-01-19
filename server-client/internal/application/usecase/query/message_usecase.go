package query

import (
	"context"
	repo "server-client/internal/domain/repository/query"
	pre_dto "server-client/internal/presenter/dto/common"

	"github.com/google/uuid"
)

type MessageRepository = repo.MessageRepo
type MessageUsecase struct {
	messageRepo MessageRepository
}

func NewMessageUsecase(messageRepo MessageRepository) *MessageUsecase {
	return &MessageUsecase{
		messageRepo: messageRepo,
	}
}

func (mu *MessageUsecase) GetMessagesByRoomID(ctx context.Context, roomID uuid.UUID, page int, size int) ([]pre_dto.MessageDTO, error) {
	msg, err := mu.messageRepo.GetMessages(roomID, page, size)
	if err != nil {
		return nil, err
	}

	var messageDTOs []pre_dto.MessageDTO
	for _, m := range msg {
		messageDTOs = append(messageDTOs, *pre_dto.NewMessageDTO(
			m.Messageid(),
			m.Roomid(),
			m.Senderid(),
			m.Content(),
			m.Isread(),
		))
	}
	return messageDTOs, nil
}
