package main

import (
	"database/sql"
	uc_query "server-client/internal/application/usecase/query"
	repo_query "server-client/internal/domain/repository/query"
	message_repo_impl "server-client/internal/infrastructure/database/query/message_repo_impl"
	recruit_repo_impl "server-client/internal/infrastructure/database/query/recruit_repo_impl"
	room_repo_impl "server-client/internal/infrastructure/database/query/room_repo_impl"
	user_repo_impl "server-client/internal/infrastructure/database/query/user_repo_impl"
	gql "server-client/internal/presenter/interface/gql/resolver"

	"github.com/google/wire"
)

type WireSet_test struct {
	Resolver       *gql.Resolver
	MessageUsecase *uc_query.MessageUsecase
	RoomUsecase    *uc_query.RoomUsecase
	UserUsecase    *uc_query.UserUsecase
	RecruitUsecase *uc_query.RecruitUsecase

	MessageRepo repo_query.MessageRepo
	RoomRepo    repo_query.RoomRepo
	UserRepo    repo_query.UserRepo
	RecruitRepo repo_query.RecruitRepo
}

func InitializeApp_test(db *sql.DB) *WireSet_test {
	wire.Build(
		gql.NewResolver,
		uc_query.NewMessageUsecase,
		uc_query.NewRoomUsecase,
		uc_query.NewUserUsecase,
		uc_query.NewRecruitUsecase,

		message_repo_impl.NewMessageRepoImpl,
		room_repo_impl.NewRoomRepoImpl,
		user_repo_impl.NewUserRepoImpl,
		recruit_repo_impl.NewRecruitRepoImpl,

		wire.Bind(new(repo_query.MessageRepo), new(*message_repo_impl.MessageRepoImpl)),
		wire.Bind(new(repo_query.RoomRepo), new(*room_repo_impl.RoomRepoImpl)),
		wire.Bind(new(repo_query.UserRepo), new(*user_repo_impl.UserRepoImpl)),
		wire.Bind(new(repo_query.RecruitRepo), new(*recruit_repo_impl.RecruitRepoImpl)),
		wire.Struct(new(WireSet_test), "*"),
	)
	return &WireSet_test{}
}
