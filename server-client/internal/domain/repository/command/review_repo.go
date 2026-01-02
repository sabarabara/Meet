package command

import "server-client/internal/application/dto"

type ReviewRepo interface {
	InsertReview(dto dto.ReviewDTO) error
}
