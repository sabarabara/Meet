# Meet

<img width="749" height="410" alt="logo1" src="https://github.com/user-attachments/assets/6abc07d3-500c-48e1-ba0c-a2ae3e56fc41" />

<img width="905" height="616" alt="スクリーンショット 2026-01-19 19 17 29" src="https://github.com/user-attachments/assets/e5e76ef0-eb58-4b3d-bd33-8dba77b96819" />

## 概要
インスタにあげる写真が欲しい人がみんなで集まって写真を撮るアプリ

新アプリ「Meet」コンセプト概要
1. 背景：現代大学生の価値判断基準
他者依存の自己定義: 際立った才能や外見を持たない限り、大学生の価値は「どれだけ友達が多いか（人脈があるか）」という外部要素で判断される傾向にある。

インスタグラムによる虚像の維持: 自身の価値を証明するため、大量の友人と過ごす様子をSNSへ投稿し続ける必要がある。それを見た側も、その「繋がり」を見て相手を評価するという循環が生まれている。

2. 課題：演出コストの高さ
「充実している自分」を演出するためのエキストラが必要だが、既存の人間レンタルサービスは費用が高く、大学生が日常的に利用するにはハードルが高い。

3. ソリューション：価値の相互扶助（物々交換）
「擬似的な関係性」の提供: お互いが「レンタルフレンド・彼氏・彼女」になりきり、撮影に協力し合う。

コストゼロの相互利用: 金銭のやり取りではなく、「自分の時間と演技力」を差し出す代わりに「自分の価値を上げる写真」を得るという、ギブ・アンド・テイクの仕組み。

相互評価システム: 撮影技術や「役者」としての振る舞いをお互いに評価し合い、演出のクオリティを担保する。

結論
「承認欲求を満たすためのコストを、ユーザー同士の『役割演じ』によって相殺し、効率的に自己演出を行うためのマッチングプラットフォーム」

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

### TLS設定
以下のコマンドをtraefik内で実行する。
```
touch cert.pem
touch key.pem

openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -sha256 -days 365 -nodes -subj '/CN=localhost'
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
│   ├── presenter
│   │   ├── middleware
│   │   │   └── auth_middleware.go
│   │   ├── interface
│   │   │   ├── gql
│   │   │   │   └── resolver
│   │   │   │       ├── messageResolver.go
│   │   │   │       ├── recruitResolver.go
│   │   │   │       ├── roomResolver.go
│   │   │   │       └── userResolver.go
│   │   │   └── rest
│   │   │
│   │   └── dto
│   │       ├── common
│   │       ├── gql
│   │       │   ├── gqlgen.yml
│   │       │   ├── graph
│   │       │   │   ├── mutation
│   │       │   │   │   ├── recruit_mutation.go
│   │       │   │   │   ├── room_mutation.go
│   │       │   │   │   └── user_mutation.go
│   │       │   │   ├── query
│   │       │   │   │   ├── message_query.go
│   │       │   │   │   ├── recruit_query.go
│   │       │   │   │   ├── room_query.go
│   │       │   │   │   └── user_query.go
│   │       │   │   └── schema.graphqls
│   │       │   └── rest
│   ├── application
│   │   └── dto
│   │       ├── message_dto.go
│   │       ├── recruit_dto.go
│   │       ├── recruit_location_dto.go
│   │       ├── review_dto.go
│   │       ├── room_dto.go
│   │       └── user_dto.go
│   ├── domain
│   │   ├── repository
│   │   │   ├── command
│   │   │   │   ├── recruit_location_repo.go
│   │   │   │   ├── recruit_repo.go
│   │   │   │   ├── review_repo.go
│   │   │   │   ├── room_repo.go
│   │   │   │   └── user_repo.go
│   │   │   └── query
│   │   │       ├── message_repo.go
│   │   │       ├── recruit_repo.go
│   │   │       ├── room_repo.go
│   │   │       └── user_repo.go
│   │   └── vo_and_entity
│   │       ├── message
│   │       │   ├── message.go
│   │       │   └── message_test.go
│   │       ├── recruit
│   │       │   ├── recruit.go
│   │       │   ├── recruit_test.go
│   │       │   └── recruit_location
│   │       │       ├── recruit_location.go
│   │       │       └── recruit_location_test.go
│   │       ├── review
│   │       │   ├── review.go
│   │       │   └── review_test.go
│   │       ├── room
│   │       │   ├── room.go
│   │       │   └── room_test.go
│   │       └── user
│   │           ├── user.go
│   │           └── user_test.go
│   └── infrastructure
│       ├── auth
│       │   ├── oidc.go
│       │   ├── login_handler.go
│       │   ├── logout_handler.go
│       │   └── session.go
│       └── database
│           ├── command
│           │   ├── recruit_location_repo_impl
│           │   │   ├── recruit_location_repo_impl.go
│           │   │   └── recruit_location_repo_impl_test.go
│           │   ├── recruit_repo_impl
│           │   │   ├── recruit_repo_impl.go
│           │   │   └── recruit_repo_impl_test.go
│           │   ├── review_repo_impl
│           │   │   ├── review_repo_impl.go
│           │   │   └── review_repo_impl_test.go
│           │   ├── room_repo_impl
│           │   │   ├── room_repo_impl.go
│           │   │   └── room_repo_impl_test.go
│           │   └── user_repo_impl
│           │       ├── user_repo_impl.go
│           │       └── user_repo_impl_test.go
│           └── query
│               ├── message_repo_impl
│               │   ├── message_repo_impl.go
│               │   └── message_repo_impl_test.go
│               ├── recruit_repo_impl
│               │   ├── recruit_repo_impl.go
│               │   └── recruit_repo_impl_test.go
│               ├── room_repo_impl
│               │   ├── room_repo_impl.go
│               │   └── room_repo_impl_test.go
│               └── user_repo_impl
│                   ├── user_repo_impl.go
│                   └── user_repo_impl_test.go
└── pkg
    └── db
        ├── dbconfig.go
        └── redisconfig.go
```
