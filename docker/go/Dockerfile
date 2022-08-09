# 2022/08/10最新versionを取得
FROM golang:1.18-alpine3.15
# アップデートとgitのインストール
RUN apk update && apk add git
# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app

# ホットリロード機能
RUN go install github.com/cosmtrek/air@v1.27.3
