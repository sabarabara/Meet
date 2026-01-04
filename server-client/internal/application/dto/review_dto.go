package dto

import (
	"github.com/google/uuid"
)

type ReviewDTO struct {
	reviewid       *uuid.UUID
	roomid         uuid.UUID
	writer         uuid.UUID
	evaluated_user uuid.UUID
	comment        string
}

func NewReviewDTO(
	reviewid *uuid.UUID,
	roomid uuid.UUID,
	writer uuid.UUID,
	evaluated_user uuid.UUID,
	comment string,
) ReviewDTO {
	return ReviewDTO{
		roomid:         roomid,
		writer:         writer,
		evaluated_user: evaluated_user,
		comment:        comment,
	}
}
func (r ReviewDTO) Reviewid() *uuid.UUID {
	return r.reviewid
}
func (r ReviewDTO) Roomid() uuid.UUID {
	return r.roomid
}
func (r ReviewDTO) Writer() uuid.UUID {
	return r.writer
}
func (r ReviewDTO) EvaluatedUser() uuid.UUID {
	return r.evaluated_user
}
func (r ReviewDTO) Comment() string {
	return r.comment
}
