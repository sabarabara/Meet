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

### 本番環境デモ実行
```
docker compose up --build
```