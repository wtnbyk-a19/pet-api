# ベースになるDocckerイメージの指定
FROM golang:1.15.6 as builder

WORKDIR /go/api

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

WORKDIR /go/api/src

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/api/bin/main \
    -ldflags '-s -w'

FROM scratch as runner

COPY --from=builder /go/api/bin/main /app/main

ENTRYPOINT ["/app/main"]