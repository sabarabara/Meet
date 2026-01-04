package dto

import "github.com/google/uuid"

type MessageDTO struct {
	messageid uuid.UUID
	roomid    uuid.UUID
	senderid  uuid.UUID
	content   string
	isread    bool
}

func NewMessageDTO(messageid, roomid, senderid uuid.UUID, content string) MessageDTO {
	return MessageDTO{
		messageid: messageid,
		roomid:    roomid,
		senderid:  senderid,
		content:   content,
		isread:    false,
	}
}

func (m MessageDTO) Messageid() uuid.UUID {
	return m.messageid
}
func (m MessageDTO) Roomid() uuid.UUID {
	return m.roomid
}
func (m MessageDTO) Senderid() uuid.UUID {
	return m.senderid
}
func (m MessageDTO) Content() string {
	return m.content
}
func (m MessageDTO) Isread() bool {
	return m.isread
}
