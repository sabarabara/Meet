package messagerepoimpl

import (
	"database/sql"
	"log"
	"server-client/internal/application/dto"
	"server-client/internal/domain/repository/query"

	"github.com/google/uuid"
)

var _ query.MessageRepo = (*MessageRepoImpl)(nil)

type messageEntity struct {
	MessageID uuid.UUID `db:"messageid"`
	RoomID    uuid.UUID `db:"roomid"`
	SenderID  uuid.UUID `db:"senderid"`
	Content   string    `db:"content"`
	Isread    bool      `db:"isread"`
}

type MessageRepoImpl struct {
	DB *sql.DB
}

func NewMessageRepoImpl(db *sql.DB) *MessageRepoImpl {
	return &MessageRepoImpl{
		DB: db,
	}
}

func (m *MessageRepoImpl) GetMessages(roomid uuid.UUID, page int, size int) ([]dto.MessageDTO, error) {
	offset := (page - 1) * size
	rows, err := m.DB.Query("SELECT messageid, roomid, senderid, content, isread FROM messages WHERE roomid = $1 LIMIT $2 OFFSET $3", roomid, size, offset)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("rows.Close error: %v", err)
			return
		}
	}()

	var messages []dto.MessageDTO
	for rows.Next() {
		var entity messageEntity
		if err := rows.Scan(&entity.MessageID, &entity.RoomID, &entity.SenderID, &entity.Content, &entity.Isread); err != nil {
			return nil, err
		}
		messageDTO := dto.NewMessageDTO(
			entity.MessageID,
			entity.RoomID,
			entity.SenderID,
			entity.Content,
		)
		messages = append(messages, messageDTO)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}
