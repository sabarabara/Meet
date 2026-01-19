package userrepoimpl

import (
	"database/sql"
	"fmt"
	"server-client/internal/application/dto"
	"server-client/internal/domain/repository/query"

	"github.com/google/uuid"
)

var _ query.UserRepo = (*UserRepoImpl)(nil)

type userentity struct {
	UserID        uuid.UUID `db:"userid"`
	UserName      string    `db:"username"`
	ImgURL        string    `db:"imgurl"`
	Pronunciation string    `db:"pronunciation"`
	SelfIntroduce string    `db:"selfintroduce"`
	Stars         float32   `db:"stars"`
}

type UserRepoImpl struct {
	DB *sql.DB
}

func NewUserRepoImpl(db *sql.DB) *UserRepoImpl {
	return &UserRepoImpl{
		DB: db,
	}
}

func (u *UserRepoImpl) GetUserByID(userID uuid.UUID) (dto.UserDTO, error) {
	var entity userentity
	err := u.DB.QueryRow(
		"SELECT userid, username, imgurl, pronunciation, selfintroduce, stars FROM users WHERE userid = $1",
		userID,
	).Scan(
		&entity.UserID,
		&entity.UserName,
		&entity.ImgURL,
		&entity.Pronunciation,
		&entity.SelfIntroduce,
		&entity.Stars,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return dto.UserDTO{}, fmt.Errorf("user not found: %w", err)
		}
		return dto.UserDTO{}, err
	}

	userDTO := dto.NewUserDTO(
		&entity.UserID,
		entity.UserName,
		entity.ImgURL,
		entity.Pronunciation,
		entity.SelfIntroduce,
		entity.Stars,
	)

	return userDTO, nil
}
