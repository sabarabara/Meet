package recruitlocationrepoimpl

import (
	"server-client/internal/application/dto"
	repo "server-client/internal/domain/repository/command"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ repo.RecruitLocationRepo = (*RecruitLocationRepoImpl)(nil)

type RecruitLocationEntity struct {
	RecruitLocationID uuid.UUID `gorm:"column:locationid;type:uuid;primaryKey;default:uuid_generate_v4()"`
	RecruitID         uuid.UUID `gorm:"column:recruitid;type:uuid;not null"`
	Latitude          float64   `gorm:"column:latitude;type:double precision;not null"`
	Longitude         float64   `gorm:"column:longitude;type:double precision;not null"`
}

type RecruitLocationRepoImpl struct {
	db *gorm.DB
}

func NewRecruitLocationRepoImpl(db *gorm.DB) *RecruitLocationRepoImpl {
	return &RecruitLocationRepoImpl{db: db}
}

func (r *RecruitLocationRepoImpl) InsertRecruitLocation(dto dto.RecruitLocationDTO) error {
	entity := RecruitLocationEntity{
		RecruitID: dto.Recruitid(),
		Latitude:  dto.Latitude(),
		Longitude: dto.Longitude(),
	}

	if err := r.db.Table("recruit_location").Create(&entity).Error; err != nil {
		return err
	}
	return nil
}
