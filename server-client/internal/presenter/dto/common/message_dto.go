package common

import "github.com/google/uuid"

type MessageDTO struct {
	msgid    uuid.UUID
	roomid   uuid.UUID
	senderid uuid.UUID
	content  string
	isread   bool
}

func NewMessageDTO(id, roomID, senderID uuid.UUID, content string, isRead bool) *MessageDTO {
	return &MessageDTO{
		msgid:    id,
		roomid:   roomID,
		senderid: senderID,
		content:  content,
		isread:   isRead,
	}
}

func (m *MessageDTO) MsgId() uuid.UUID {
	return m.msgid
}

func (m *MessageDTO) RoomId() uuid.UUID {
	return m.roomid
}

func (m *MessageDTO) SenderId() uuid.UUID {
	return m.senderid
}

func (m *MessageDTO) Content() string {
	return m.content
}

func (m *MessageDTO) IsRead() bool {
	return m.isread
}
