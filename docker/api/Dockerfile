# ベースになるDocckerイメージの指定
FROM golang:latest

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY ../../ ./

ARG CGO_ENABLED=0

CMD ["go", "run", "server.go"]