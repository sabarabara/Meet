package roomrepoimpl

import (
	"server-client/internal/application/dto"
	repo "server-client/internal/domain/repository/command"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ repo.RoomRepo = (*RoomRepoImpl)(nil)

type RoomEntity struct {
	Roomid    uuid.UUID `gorm:"column:roomid;type:uuid;primaryKey;default:uuid_generate_v4()"`
	Recruitid uuid.UUID `gorm:"column:recruitid;type:uuid"`
	Userid    uuid.UUID `gorm:"column:userid;type:uuid"`
	Role      string    `gorm:"column:role;type:varchar(100)"`
	Isfinish  bool      `gorm:"column:isfinish;type:boolean"`
}

type RoomRepoImpl struct {
	db *gorm.DB
}

func NewRoomRepoImpl(db *gorm.DB) *RoomRepoImpl {
	return &RoomRepoImpl{db: db}
}

func (r *RoomRepoImpl) InsertRoom(dto dto.RoomDTO) error {
	entity := RoomEntity{
		Recruitid: dto.Recruitid(),
		Userid:    dto.Userid(),
		Role:      dto.Role(),
		Isfinish:  dto.Isfinished(),
	}
	if err := r.db.Table("rooms").Create(&entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *RoomRepoImpl) DeleteRoom(roomid uuid.UUID) error {
	if err := r.db.Table("rooms").Where("roomid = ?", roomid).Delete(&RoomEntity{}).Error; err != nil {
		return err
	}
	return nil
}
