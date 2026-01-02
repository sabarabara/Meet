package recruitrepoimpl

import (
	"server-client/internal/application/dto"
	repo "server-client/internal/domain/repository/command"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ repo.RecruitRepo = (*RecruitRepoImpl)(nil)

type RecruitEntity struct {
	Recruitid    uuid.UUID `gorm:"column:recruitid;type:uuid;primaryKey;default:uuid_generate_v4()"`
	Userid       uuid.UUID `gorm:"column:userid;type:uuid;not null"`
	Area         string    `gorm:"column:area;type:varchar(255);not null"`
	Imgurl       string    `gorm:"column:imgurl;type:varchar(255);not null"`
	Man          uint16    `gorm:"column:man;type:smallint;not null"`
	Woman        uint16    `gorm:"column:woman;type:smallint;not null"`
	Vacant_man   uint16    `gorm:"column:vacant_man;type:smallint;not null"`
	Vacant_woman uint16    `gorm:"column:vacant_woman;type:smallint;not null"`
	Comment      string    `gorm:"column:comment;type:text;not null"`
	Date         time.Time `gorm:"column:date;type:timestamp;not null"`
}

type RecruitRepoImpl struct {
	db *gorm.DB
}

func NewRecruitRepoImpl(db *gorm.DB) *RecruitRepoImpl {
	return &RecruitRepoImpl{db: db}
}

func (r *RecruitRepoImpl) InsertRecruit(dto dto.RecruitDTO) error {
	entity := RecruitEntity{
		Userid:       dto.Userid(),
		Area:         dto.Area(),
		Imgurl:       dto.Imgurl(),
		Man:          dto.Man(),
		Woman:        dto.Woman(),
		Vacant_man:   dto.VacantMan(),
		Vacant_woman: dto.VacantWoman(),
		Comment:      dto.Commnt(),
		Date:         dto.Date(),
	}

	if err := r.db.Table("recruit").Create(&entity).Error; err != nil {
		return err
	}
	return nil
}
