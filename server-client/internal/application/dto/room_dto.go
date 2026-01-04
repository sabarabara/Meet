package dto

import "github.com/google/uuid"

type RoomDTO struct {
	recruitid  uuid.UUID
	userid     uuid.UUID
	role       string
	isfinished bool
}

func NewRoomDTO(
	recruitid uuid.UUID,
	userid uuid.UUID,
	role string,
	isfinished bool,
) RoomDTO {
	return RoomDTO{
		recruitid:  recruitid,
		userid:     userid,
		role:       role,
		isfinished: isfinished,
	}
}

func (r RoomDTO) Recruitid() uuid.UUID {
	return r.recruitid
}
func (r RoomDTO) Userid() uuid.UUID {
	return r.userid
}
func (r RoomDTO) Role() string {
	return r.role
}
func (r RoomDTO) Isfinished() bool {
	return r.isfinished
}
