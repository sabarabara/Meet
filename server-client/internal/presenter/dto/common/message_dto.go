package common

type MessageDTO struct {
	msgid    string
	roomid   string
	senderid string
	content  string
	isread   bool
}

func NewMessageDTO(id, roomID, senderID, content string, isRead bool) *MessageDTO {
	return &MessageDTO{
		msgid:    id,
		roomid:   roomID,
		senderid: senderID,
		content:  content,
		isread:   isRead,
	}
}

func (m *MessageDTO) MsgId() string {
	return m.msgid
}

func (m *MessageDTO) RoomId() string {
	return m.roomid
}

func (m *MessageDTO) SenderId() string {
	return m.senderid
}

func (m *MessageDTO) Content() string {
	return m.content
}

func (m *MessageDTO) IsRead() bool {
	return m.isread
}
