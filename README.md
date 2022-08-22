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

## 採用技術補足

- [Gin Web Framework](https://gin-gonic.com/ja/docs/)
- [Air](https://github.com/cosmtrek/air) ※`yarn run watch`的なホットリロード機能