package command

import (
	"server-client/internal/application/dto"
)

type RecruitRepo interface {
	InsertRecruit(dto dto.RecruitDTO) error
}
