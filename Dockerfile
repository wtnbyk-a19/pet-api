# ベースになるDocckerイメージの指定
FROM golang:latest
# コンテナ内にディレクトリを作成
RUN mkdir /app
# コンテナログイン時のディレクトリを指定
WORKDIR /app