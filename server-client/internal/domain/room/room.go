package room

import (
	"errors"

	"github.com/google/uuid"
)

const (
	errorEmptyRoomID    = "room ID cannot be empty"
	errorEmptyRecruitID = "recruit ID cannot be empty"
	errorEmptyUserID    = "user ID cannot be empty"
	errorEmptyRole      = "role cannot be empty"
)

type Room struct {
	roomid     uuid.UUID
	recruitid  uuid.UUID
	userid     uuid.UUID
	role       string
	isfinished bool
}

func NewRoom(roomid uuid.UUID, recruitid uuid.UUID, userid uuid.UUID, role string, isfinished bool) (Room, error) {
	if roomid == uuid.Nil {
		return Room{}, errors.New(errorEmptyRoomID)
	}
	if recruitid == uuid.Nil {
		return Room{}, errors.New(errorEmptyRecruitID)
	}
	if userid == uuid.Nil {
		return Room{}, errors.New(errorEmptyUserID)
	}
	if role == "" {
		return Room{}, errors.New(errorEmptyRole)
	}

	return Room{
		roomid:     roomid,
		recruitid:  recruitid,
		userid:     userid,
		role:       role,
		isfinished: isfinished,
	}, nil
}

func (r Room) Roomid() uuid.UUID {
	return r.roomid
}
func (r Room) Recruitid() uuid.UUID {
	return r.recruitid
}
func (r Room) Userid() uuid.UUID {
	return r.userid
}
func (r Room) Role() string {
	return r.role
}
func (r Room) Isfinished() bool {
	return r.isfinished
}
