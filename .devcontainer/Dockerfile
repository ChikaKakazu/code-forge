FROM golang:1.23.4 AS dev

RUN apt-get update && apt-get install -y \
    && rm -rf /var/lib/apt/lists/*

# 作業ディレクトリ作成
WORKDIR /workspace

# 環境変数など（必要なら）
ENV CGO_ENABLED=0
ENV GO111MODULE=on

# 必要なGoツールやライブラリをインストール
RUN go install github.com/spf13/cobra-cli@latest

# ホスト側でプロジェクトをマウントする前提
# go.mod, go.sum, srcファイルはコンテナ起動後にローカルからマウントされる

