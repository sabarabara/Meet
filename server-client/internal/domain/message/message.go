package message

import (
	"errors"

	"github.com/google/uuid"
)

const (
	errorEmptyMessageID = "メッセージIDが空になっています"
	errorEmptyRoomID    = "ルームIDが空になっています"
	errorEmptySenderID  = "送信者IDが空になっています"
	errorEmptyContent   = "メッセージ内容が空になっています"
)

type Message struct {
	messageid uuid.UUID
	roomid    uuid.UUID
	senderid  uuid.UUID
	content   string
	isread    bool
}

func NewMessage(messageid uuid.UUID, roomid uuid.UUID, senderid uuid.UUID, content string, isread bool) (Message, error) {
	if messageid == uuid.Nil {
		return Message{}, errors.New(errorEmptyMessageID)
	}
	if roomid == uuid.Nil {
		return Message{}, errors.New(errorEmptyRoomID)
	}
	if senderid == uuid.Nil {
		return Message{}, errors.New(errorEmptySenderID)
	}
	if content == "" {
		return Message{}, errors.New(errorEmptyContent)
	}
	return Message{
		messageid: messageid,
		roomid:    roomid,
		senderid:  senderid,
		content:   content,
		isread:    isread,
	}, nil
}

func (m Message) Messageid() uuid.UUID {
	return m.messageid
}
func (m Message) Roomid() uuid.UUID {
	return m.roomid
}
func (m Message) Senderid() uuid.UUID {
	return m.senderid
}
func (m Message) Content() string {
	return m.content
}
func (m Message) Isread() bool {
	return m.isread
}
