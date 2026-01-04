package userrepoimpl

import (
	"server-client/internal/application/dto"
	repo "server-client/internal/domain/repository/command"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ repo.UserRepo = (*UserRepoImpl)(nil)

type UserEntity struct {
	Userid        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username      string    `gorm:"type:varchar(100);not null"`
	Imgurl        string    `gorm:"type:varchar(255)"`
	Pronunciation string    `gorm:"type:varchar(255)"`
	Selfintroduce string    `gorm:"type:text"`
	Stars         float32   `gorm:"type:float;default:0"`
}

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepoImpl(db *gorm.DB) *UserRepoImpl {
	return &UserRepoImpl{db: db}
}

func (r *UserRepoImpl) UpsertUser(d dto.UserDTO) (dto.UserDTO, error) {
	if d.Userid() != nil {
		if err := r.db.Table("users").
			Where("userid = ?", *d.Userid()).
			Updates(map[string]interface{}{
				"username":      d.Username(),
				"imgurl":        d.Imgurl(),
				"pronunciation": d.Pronunciation(),
				"selfintroduce": d.Selfintroduction(),
				"stars":         d.Stars(),
			}).Error; err != nil {
			return dto.UserDTO{}, err
		}

		return d, nil
	}

	entity := UserEntity{
		Username:      d.Username(),
		Imgurl:        d.Imgurl(),
		Pronunciation: d.Pronunciation(),
		Selfintroduce: d.Selfintroduction(),
		Stars:         d.Stars(),
	}

	if err := r.db.Table("users").Create(&entity).Error; err != nil {
		return dto.UserDTO{}, err
	}

	return dto.NewUserDTO(
		&entity.Userid,
		entity.Username,
		entity.Imgurl,
		entity.Pronunciation,
		entity.Selfintroduce,
		entity.Stars,
	), nil
}

func (r *UserRepoImpl) DeleteUser(userid uuid.UUID) error {
	if err := r.db.Table("users").Where("userid = ?", userid).Delete(&UserEntity{}).Error; err != nil {
		return err
	}
	return nil
}
