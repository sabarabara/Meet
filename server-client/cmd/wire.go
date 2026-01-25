package main

import (
	"context"
	"database/sql"
	uc_query "server-client/internal/application/usecase/query"
	repo_command "server-client/internal/domain/repository/command"
	repo_query "server-client/internal/domain/repository/query"
	"server-client/internal/infrastructure/auth"
	oidc "server-client/internal/infrastructure/auth/oidc_client"
	strategy "server-client/internal/infrastructure/auth/provider_strategy"
	"server-client/internal/infrastructure/auth/provider_strategy/provider"
	auth_command_repo_impl "server-client/internal/infrastructure/database/command/auth_repo_impl"
	command_user_repo_impl "server-client/internal/infrastructure/database/command/user_repo_impl"
	auth_query_repo_impl "server-client/internal/infrastructure/database/query/auth_repo_impl"
	message_repo_impl "server-client/internal/infrastructure/database/query/message_repo_impl"
	recruit_repo_impl "server-client/internal/infrastructure/database/query/recruit_repo_impl"
	room_repo_impl "server-client/internal/infrastructure/database/query/room_repo_impl"
	query_user_repo_impl "server-client/internal/infrastructure/database/query/user_repo_impl"
	gql "server-client/internal/presenter/interface/gql/resolver"
	pkg_auth "server-client/pkg/auth"
	pkg_redis "server-client/pkg/db"

	"github.com/google/wire"
	"gorm.io/gorm"
)

type WireSet_Test struct {
	context.Context
	Resolver       *gql.Resolver
	MessageUsecase *uc_query.MessageUsecase
	RoomUsecase    *uc_query.RoomUsecase
	UserUsecase    *uc_query.UserUsecase
	RecruitUsecase *uc_query.RecruitUsecase

	MessageQueryRepo repo_query.MessageRepo
	RoomQueryRepo    repo_query.RoomRepo
	UserQueryRepo    repo_query.UserRepo
	RecruitQueryRepo repo_query.RecruitRepo
	AuthQueryRepo    repo_query.AuthRepo

	AuthCommandRepo repo_command.AuthRepo
	UserCommandRepo repo_command.UserRepo

	LoginHandler     *auth.LoginHandler
	SessionManager   *auth.SessionManager
	RedisClient      *pkg_redis.RedisClient
	ProviderRegistry *strategy.ProviderRegistry
	OIDCService      *oidc.OIDCService
	GoogleProvider   *provider.GoogleStrategy
}

func InitializeApp_Test(sqlDB *sql.DB, gormDB *gorm.DB, redisClient *pkg_redis.RedisClient) (*WireSet_Test, error) {
	wire.Build(
		//外部パッケージ
		context.Background,
		pkg_auth.LoadAuthConfigs,

		//interface層
		gql.NewResolver,

		//application層
		uc_query.NewMessageUsecase,
		uc_query.NewRoomUsecase,
		uc_query.NewUserUsecase,
		uc_query.NewRecruitUsecase,

		//infrastructure層
		//repository - query
		message_repo_impl.NewMessageRepoImpl,
		room_repo_impl.NewRoomRepoImpl,
		query_user_repo_impl.NewUserRepoImpl,
		recruit_repo_impl.NewRecruitRepoImpl,
		auth_query_repo_impl.NewAuthRepoImpl,

		//repository - command
		auth_command_repo_impl.NewAuthRepoImpl,
		command_user_repo_impl.NewUserRepoImpl,

		//ここに新しいプロバイダのStrategyを追加してね!!
		provider.NewGoogleStrategy,

		//infrastructure - auth
		strategy.NewProviderRegistry,
		oidc.NewOIDCService,
		auth.NewLoginHandler,
		auth.NewSessionManager,

		//ここ新しく追加したプロバイダを単体として抽出してね!!
		wire.FieldsOf(new(*pkg_auth.AuthConfigs), "Google"),

		//interfaceを使ってDIする時はここでバインドしてね!!
		wire.Bind(new(repo_query.MessageRepo), new(*message_repo_impl.MessageRepoImpl)),
		wire.Bind(new(repo_query.RoomRepo), new(*room_repo_impl.RoomRepoImpl)),
		wire.Bind(new(repo_query.UserRepo), new(*query_user_repo_impl.UserRepoImpl)),
		wire.Bind(new(repo_query.RecruitRepo), new(*recruit_repo_impl.RecruitRepoImpl)),
		wire.Bind(new(repo_query.AuthRepo), new(*auth_query_repo_impl.AuthRepoImpl)),
		wire.Bind(new(repo_command.AuthRepo), new(*auth_command_repo_impl.AuthRepoImpl)),
		wire.Bind(new(repo_command.UserRepo), new(*command_user_repo_impl.UserRepoImpl)),

		wire.Struct(new(WireSet_Test), "*"),
	)
	return &WireSet_Test{}, nil
}
