package query

import (
	"context"
	repo "server-client/internal/domain/repository/query"
	pre_dto "server-client/internal/presenter/dto/common"
)

type RecruitRepository = repo.RecruitRepo
type RecruitUsecase struct {
	recruitrepo RecruitRepository
}

func NewRecruitUsecase(recruitrepo RecruitRepository) *RecruitUsecase {
	return &RecruitUsecase{
		recruitrepo: recruitrepo,
	}
}

func (ru *RecruitUsecase) GetRecruit(ctx context.Context, page int, size int) ([]pre_dto.RecruitDTO, error) {
	recruit, recruit_location, err := ru.recruitrepo.GetRecruits(page, size)
	if err != nil {
		return nil, err
	}

	if len(recruit) != len(recruit_location) {
		return nil, nil
	}

	var recruitDTOs []pre_dto.RecruitDTO
	for i := range len(recruit) {
		recruitDTOs = append(recruitDTOs, *pre_dto.NewRecruitDTO(
			*recruit[i].Recruitid(),
			recruit[i].Userid(),
			recruit[i].Area(),
			recruit[i].Imgurl(),
			recruit[i].Man(),
			recruit[i].Woman(),
			recruit[i].VacantMan(),
			recruit[i].VacantWoman(),
			recruit[i].Commnt(),
			recruit[i].Date(),
			recruit_location[i].Latitude(),
			recruit_location[i].Longitude(),
		))
	}
	return recruitDTOs, nil
}
