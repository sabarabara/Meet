package userrepoimpl

import (
	"database/sql"
	"log"
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

func (u *UserRepoImpl) GetUsers(page int, size int) ([]dto.UserDTO, error) {
	offset := (page - 1) * size
	rows, err := u.DB.Query("SELECT userid, username, imgurl, pronunciation, selfintroduce, stars FROM users LIMIT $1 OFFSET $2", size, offset)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("rows.Close error: %v", err)
			return
		}
	}()

	var users []dto.UserDTO
	for rows.Next() {
		var entity userentity
		if err := rows.Scan(&entity.UserID, &entity.UserName, &entity.ImgURL, &entity.Pronunciation, &entity.SelfIntroduce, &entity.Stars); err != nil {
			return nil, err
		}
		userDTO := dto.NewUserDTO(
			&entity.UserID,
			entity.UserName,
			entity.ImgURL,
			entity.Pronunciation,
			entity.SelfIntroduce,
			entity.Stars,
		)
		users = append(users, userDTO)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
