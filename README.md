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

## デバッグ設定

- 任意の箇所にブレイクポイントを設定する
- アプリケーションファイルの場合は`main.go`あるいは、テストファイルの場合は`main_test.go`を開く
- 上記ファイルのルートメソッドの"▶︎"ボタンをクリック

設定中にエラーが発生したら[こちら](https://medium.com/@gorlemkun/goland%E3%81%A7%E3%83%96%E3%83%AC%E3%83%BC%E3%82%AF%E3%83%9D%E3%82%A4%E3%83%B3%E3%83%88%E3%83%87%E3%83%90%E3%83%83%E3%82%B0%E3%81%8C%E5%87%BA%E6%9D%A5%E3%82%8B%E3%82%88%E3%81%86%E3%81%AB%E3%81%AA%E3%82%8B%E3%81%BE%E3%81%A7-8f0a63cd3804)

## ライブラリ補足

- [Gin Web Framework](https://gin-gonic.com/ja/docs/)
  - 軽量webフレームワーク
  - 本アプリではルーティング設定で利用

- [Air](https://github.com/cosmtrek/air)
  - ホットリロード機能
  - フロント開発でいう`yarn run watch`的なもの