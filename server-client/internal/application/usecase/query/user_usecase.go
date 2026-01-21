package query

import (
	"context"
	repo "server-client/internal/domain/repository/query"
	pre_dto "server-client/internal/presenter/dto/common"

	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepo repo.UserRepo
}

func NewUserUsecase(userRepo repo.UserRepo) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}
func (uu *UserUsecase) GetUsers(ctx context.Context, userid uuid.UUID) (*pre_dto.UserDTO, error) {
	user, err := uu.userRepo.GetUserByID(userid)
	if err != nil {
		return nil, err
	}

	userDTO := pre_dto.NewUserDTO(
		*user.Userid(),
		user.Username(),
		user.Stars(),
		user.Imgurl(),
		user.Pronunciation(),
		user.Selfintroduction(),
	)
	return userDTO, nil
}
