package query

import (
	"context"
	repo "server-client/internal/domain/repository/query"
	pre_dto "server-client/internal/presenter/dto/common"

	"github.com/google/uuid"
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

func (ru *RoomUsecase) GetRoomsByUserID(ctx context.Context, userID uuid.UUID, page int, size int) ([]pre_dto.RoomDTO, error) {
	rooms, err := ru.roomRepo.GetRooms(page, size)
	if err != nil {
		return nil, err
	}

	var roomDTOs []pre_dto.RoomDTO
	for _, room := range rooms {
		roomDTOs = append(roomDTOs, *pre_dto.NewRoomDTO(
			*room.Roomid(),
			room.Recruitid(),
			room.Userid(),
			room.Role(),
			room.Isfinished(),
		))
	}
	return roomDTOs, nil
}
