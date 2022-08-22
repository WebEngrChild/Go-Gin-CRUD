## 環境構築

```
# 環境変数設定
$ cp .env.example .env

# docker環境立ち上げ
$ docker compose build
$ docker compose up -d

# アプリケーション起動
$ go run main.go

# テスト実行
$ go test
```