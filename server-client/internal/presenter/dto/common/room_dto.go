package common

type RoomDTO struct {
	roomID     string
	recruitID  string
	userID     string
	role       string
	isFinished bool
}

func NewRoomDTO(roomID, recruitID, userID, role string, isFinished bool) *RoomDTO {
	return &RoomDTO{
		roomID:     roomID,
		recruitID:  recruitID,
		userID:     userID,
		role:       role,
		isFinished: isFinished,
	}
}

func (r *RoomDTO) RoomID() string {
	return r.roomID
}

func (r *RoomDTO) RecruitID() string {
	return r.recruitID
}

func (r *RoomDTO) UserID() string {
	return r.userID
}

func (r *RoomDTO) Role() string {
	return r.role
}

func (r *RoomDTO) IsFinished() bool {
	return r.isFinished
}
