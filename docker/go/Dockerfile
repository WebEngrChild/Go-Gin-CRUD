FROM golang:1.18-alpine3.15

RUN apk update && \
    apk --no-cache add git

RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app

RUN go install github.com/cosmtrek/air@v1.27.3
CMD ["air", "-c", ".air.toml"]