package resolver

import (
	"server-client/internal/application/usecase/query"
)

type Resolver struct {
	MessageUsecase *query.MessageUsecase
	RoomUsecase    *query.RoomUsecase
	UserUsecase    *query.UserUsecase
	RecruitUsecase *query.RecruitUsecase
}

func NewResolver(
	messageUsecase *query.MessageUsecase,
	roomUsecase *query.RoomUsecase,
	userUsecase *query.UserUsecase,
	recruitUsecase *query.RecruitUsecase,
) *Resolver {
	return &Resolver{
		MessageUsecase: messageUsecase,
		RoomUsecase:    roomUsecase,
		UserUsecase:    userUsecase,
		RecruitUsecase: recruitUsecase,
	}
}
