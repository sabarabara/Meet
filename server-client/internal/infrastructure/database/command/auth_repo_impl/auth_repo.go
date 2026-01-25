package authrepoimpl

import (
	repo "server-client/internal/domain/repository/command"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ repo.AuthRepo = (*AuthRepoImpl)(nil)

type AuthEntity struct {
	UserID   uuid.UUID `gorm:"column:userid;type:varchar(255);not null"`
	Sub      string    `gorm:"column:sub;type:varchar(255);not null"`
	Provider string    `gorm:"column:provider;type:varchar(255);not null"`
}

type AuthRepoImpl struct {
	db *gorm.DB
}

func NewAuthRepoImpl(db *gorm.DB) *AuthRepoImpl {
	return &AuthRepoImpl{db: db}
}

func (r *AuthRepoImpl) CreateAuthenticatedTable(userid uuid.UUID, sub string, provider string) error {
	entity := AuthEntity{
		UserID:   userid,
		Sub:      sub,
		Provider: provider,
	}

	if err := r.db.Table("authentication").Create(&entity).Error; err != nil {
		return err
	}
	return nil
}
