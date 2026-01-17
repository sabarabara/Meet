package query

import (
	repo "server-client/internal/domain/repository/query"
	pre_dto "server-client/internal/presenter/dto/common"
)

type UserUsecase struct {
	userRepo repo.UserRepo
}

func NewUserUsecase(userRepo repo.UserRepo) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}
func (uu *UserUsecase) GetUsers(page int, size int) ([]*pre_dto.UserDTO, error) {
	users, err := uu.userRepo.GetUsers(page, size)
	if err != nil {
		return nil, err
	}
	var userDTOs []*pre_dto.UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, pre_dto.NewUserDTO(
			user.Userid().String(),
			user.Username(),
			user.Stars(),
			user.Imgurl(),
			user.Pronunciation(),
			user.Selfintroduction(),
		))
	}
	return userDTOs, nil
}
