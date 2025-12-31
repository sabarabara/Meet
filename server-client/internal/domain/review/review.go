package review

import (
	"errors"

	"github.com/google/uuid"
)

const (
	errorEmptyReviewID        = "レビューIDが空になっています"
	errorEmptyRoomID          = "ルームIDが空になっています"
	errorEmptyWriterID        = "書き込み者IDが空になっています"
	errorEmptyEvaluatedUserID = "評価されるユーザーIDが空になっています"
	errorEmptyComment         = "コメントが空になっています"
)

type Review struct {
	reviewid       uuid.UUID
	roomid         uuid.UUID
	writer         uuid.UUID
	evaluated_user uuid.UUID
	comment        string
}

func NewReview(reviewid uuid.UUID, roomid uuid.UUID, writer uuid.UUID, evaluated_user uuid.UUID, comment string) (Review, error) {
	if reviewid == uuid.Nil {
		return Review{}, errors.New(errorEmptyReviewID)
	}
	if roomid == uuid.Nil {
		return Review{}, errors.New(errorEmptyRoomID)
	}
	if writer == uuid.Nil {
		return Review{}, errors.New(errorEmptyWriterID)
	}
	if evaluated_user == uuid.Nil {
		return Review{}, errors.New(errorEmptyEvaluatedUserID)
	}
	if comment == "" {
		return Review{}, errors.New(errorEmptyComment)
	}

	return Review{
		reviewid:       reviewid,
		roomid:         roomid,
		writer:         writer,
		evaluated_user: evaluated_user,
		comment:        comment,
	}, nil
}

func (r Review) Reviewid() uuid.UUID {
	return r.reviewid
}
func (r Review) Roomid() uuid.UUID {
	return r.roomid
}
func (r Review) Writer() uuid.UUID {
	return r.writer
}
func (r Review) EvaluatedUser() uuid.UUID {
	return r.evaluated_user
}
func (r Review) Comment() string {
	return r.comment
}
