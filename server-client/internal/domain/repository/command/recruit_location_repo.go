package command

import "server-client/internal/application/dto"

type RecruitLocationRepo interface {
	InsertRecruitLocation(dto dto.RecruitLocationDTO) error
}
