package query

import "server-client/internal/application/dto"

type RecruitRepo interface {
	GetRecruits(page int, size int) ([]dto.RecruitDTO, []dto.RecruitLocationDTO, error)
}
