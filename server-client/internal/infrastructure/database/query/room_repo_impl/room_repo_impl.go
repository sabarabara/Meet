package roomrepoimpl

import (
	"database/sql"
	"log"
	"server-client/internal/application/dto"
	"server-client/internal/domain/repository/query"

	"github.com/google/uuid"
)

var _ query.RoomRepo = (*RoomRepoImpl)(nil)

type roomentity struct {
	roomid     uuid.UUID `db:"roomid"`
	recruitid  uuid.UUID `db:"recruitid"`
	userid     uuid.UUID `db:"userid"`
	role       string    `db:"role"`
	isfinished bool      `db:"isfinish"`
}

type RoomRepoImpl struct {
	DB *sql.DB
}

func NewRoomRepoImpl(db *sql.DB) *RoomRepoImpl {
	return &RoomRepoImpl{
		DB: db,
	}
}

func (r *RoomRepoImpl) GetRooms(page int, size int) ([]dto.RoomDTO, error) {
	offset := (page - 1) * size
	rows, err := r.DB.Query("SELECT roomid, recruitid, userid, role, isfinish FROM rooms LIMIT $1 OFFSET $2", size, offset)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("rows.Close error: %v", err)
			return
		}
	}()

	var rooms []dto.RoomDTO
	for rows.Next() {
		var entity roomentity
		if err := rows.Scan(&entity.roomid, &entity.recruitid, &entity.userid, &entity.role, &entity.isfinished); err != nil {
			return nil, err
		}
		roomDTO := dto.NewRoomDTO(
			&entity.roomid,
			entity.recruitid,
			entity.userid,
			entity.role,
			entity.isfinished,
		)
		rooms = append(rooms, roomDTO)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return rooms, nil
}
