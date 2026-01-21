package common

import "github.com/google/uuid"

type RoomDTO struct {
	roomID     uuid.UUID
	recruitID  uuid.UUID
	userID     uuid.UUID
	role       string
	isFinished bool
}

func NewRoomDTO(roomID, recruitID, userID uuid.UUID, role string, isFinished bool) *RoomDTO {
	return &RoomDTO{
		roomID:     roomID,
		recruitID:  recruitID,
		userID:     userID,
		role:       role,
		isFinished: isFinished,
	}
}

func (r *RoomDTO) RoomID() uuid.UUID {
	return r.roomID
}

func (r *RoomDTO) RecruitID() uuid.UUID {
	return r.recruitID
}

func (r *RoomDTO) UserID() uuid.UUID {
	return r.userID
}

func (r *RoomDTO) Role() string {
	return r.role
}

func (r *RoomDTO) IsFinished() bool {
	return r.isFinished
}
