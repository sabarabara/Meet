package reviewrepoimpl

import (
	"server-client/internal/application/dto"
	repo "server-client/internal/domain/repository/command"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ repo.ReviewRepo = (*ReviewRepoImpl)(nil)

type ReviewEntity struct {
	Reviewid      uuid.UUID `gorm:"column:reviewid;type:uuid;primaryKey;default:uuid_generate_v4()"`
	Roomid        uuid.UUID `gorm:"column:roomid;type:uuid;not null"`
	Writer        uuid.UUID `gorm:"column:writer;type:uuid;not null"`
	EvaluatedUser uuid.UUID `gorm:"column:evaluated_user;type:uuid;not null"`
	Comment       string    `gorm:"column:comment;type:text;not null"`
}

type ReviewRepoImpl struct {
	db *gorm.DB
}

func NewReviewRepoImpl(db *gorm.DB) *ReviewRepoImpl {
	return &ReviewRepoImpl{db: db}
}
func (r *ReviewRepoImpl) InsertReview(dto dto.ReviewDTO) error {
	entity := ReviewEntity{
		Roomid:        dto.Roomid(),
		Writer:        dto.Writer(),
		EvaluatedUser: dto.EvaluatedUser(),
		Comment:       dto.Comment(),
	}

	if err := r.db.Table("review").Create(&entity).Error; err != nil {
		return err
	}
	return nil
}
