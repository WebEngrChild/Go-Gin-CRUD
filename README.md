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

## ライブラリ補足

- [Gin Web Framework](https://gin-gonic.com/ja/docs/)
  - 軽量webフレームワーク
  - 本アプリではルーティング設定で利用

- [Air](https://github.com/cosmtrek/air)
  - ホットリロード機能
  - フロント開発でいう`yarn run watch`的なもの