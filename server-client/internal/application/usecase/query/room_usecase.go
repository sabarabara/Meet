package query

import (
	repo "server-client/internal/domain/repository/query"
	pre_dto "server-client/internal/presenter/dto/common"
)

type RoomRepository = repo.RoomRepo

type RoomUsecase struct {
	roomRepo RoomRepository
}

func NewRoomUsecase(roomRepo RoomRepository) *RoomUsecase {
	return &RoomUsecase{
		roomRepo: roomRepo,
	}
}

func (ru *RoomUsecase) GetRoomsByUserID(userID string, page int, size int) ([]pre_dto.RoomDTO, error) {
	rooms, err := ru.roomRepo.GetRooms(page, size)
	if err != nil {
		return nil, err
	}

	var roomDTOs []pre_dto.RoomDTO
	for _, room := range rooms {
		roomDTOs = append(roomDTOs, *pre_dto.NewRoomDTO(
			room.Roomid().String(),
			room.Recruitid().String(),
			room.Userid().String(),
			room.Role(),
			room.Isfinished(),
		))
	}
	return roomDTOs, nil
}
