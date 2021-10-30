# ベースになるDocckerイメージの指定
FROM golang:latest

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY .env ./

COPY ./src/server.go  ./

ARG CGO_ENABLED=0

CMD ["go", "run", "server.go"]