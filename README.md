# Meet

## 概要
インスタにあげる写真が欲しい人がみんなで集まって写真を撮るアプリ

## 開発者向け

### バックエンド コード整形手順
下記のコマンドをgit add前に実行する
```
gofmt -s -w .
goimports -w .
golangci-lint run
```

```
export DB_USER=postgres
export DB_PASSWORD=secret
export DB_NAME=mydb
export DB_HOST=db
```

### 本番環境デモ実行
```
docker compose up --build
```

### テスト実行コマンド
**単体テスト**

プロジェクトルートで
```
go test ./...
```
**docker-compose上でテスト**
```
bash ./run-tests.sh
```

### ファイル構成(仮)
```
.
├── cmd
│   └── main.go
├── docker
│   ├── dev
│   │   └── Dockerfile
│   └── prod
├── go.mod
├── go.sum
├── internal
│   ├── application
│   │   ├── factory
│   │   └── usecase
│   │       ├── command
│   │       │   ├── recruit
│   │       │   │   ├── recruit_service.go
│   │       │   │   └── recruit_usecase.go
│   │       │   ├── room
│   │       │   │   ├── room_service.go
│   │       │   │   └── room_usecase.go
│   │       │   └── user
│   │       │       ├── user_service.go
│   │       │       └── user_usecase.go
│   │       └── query
│   │           ├── message
│   │           │   └── message_usecase.go
│   │           ├── recruit
│   │           │   └── recruit_usecase.go
│   │           ├── room
│   │           │   └── room_usecase.go
│   │           └── user
│   │               └── user_usecase.go
│   ├── di
│   │   └── wire.go
│   ├── domain
│   │   ├── message.go
│   │   ├── recruit.go
│   │   ├── recruit_location.go
│   │   ├── review.go
│   │   ├── room.go
│   │   └── user.go
│   ├── infrastructure
│   │   ├── database
│   │   │   ├── command
│   │   │   │   ├── impl
│   │   │   │   │   ├── recruit_read_repo_impl.go
│   │   │   │   │   ├── recruit_repo_impl.go
│   │   │   │   │   ├── review_repo_impl.go
│   │   │   │   │   └── user_repo_impl.go
│   │   │   │   └── repo
│   │   │   │       ├── recruit_read_repo.go
│   │   │   │       ├── recruit_repo.go
│   │   │   │       ├── review_repo.go
│   │   │   │       └── user_repo.go
│   │   │   └── query
│   │   │       ├── impl
│   │   │       │   ├── message_repo_impl.go
│   │   │       │   ├── recruit_repo_impl.go
│   │   │       │   ├── room_repo_impl.go
│   │   │       │   └── user_repo_impl.go
│   │   │       └── repo
│   │   │           ├── message_repo.go
│   │   │           ├── recruit_repo.go
│   │   │           ├── room_repo.go
│   │   │           └── user_repo.go
│   │   └── score_prediction
│   ├── presenter
│   │   ├── dto
│   │   │   ├── common
│   │   │   ├── gql
│   │   │   │   ├── gqlgen.yml
│   │   │   │   └── graph
│   │   │   │       ├── mutation
│   │   │   │       │   ├── recruit_mutation.go
│   │   │   │       │   ├── room_mutation.go
│   │   │   │       │   └── user_mutation.go
│   │   │   │       ├── query
│   │   │   │       │   ├── message_query.go
│   │   │   │       │   ├── recruit_query.go
│   │   │   │       │   ├── room_query.go
│   │   │   │       │   └── user_query.go
│   │   │   │       └── schema.graphqls
│   │   │   └── rest
│   │   └── interface
│   │       └── gql
│   │           └── resolver
│   │               ├── messageResolver.go
│   │               ├── recruitResolver.go
│   │               ├── roomResolver.go
│   │               └── userResolver.go
├── pkg
│   ├── auth
│   │   ├── login_handler.go
│   │   ├── logout_handler.go
│   │   ├── oidc.go
│   │   └── session.go
│   ├── db
│   │   ├── dbconfig.go
│   │   └── redisconfig.go
│   └── middleware
│       └── middleware.go

```