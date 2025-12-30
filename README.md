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