# ベースとなるDockerイメージ指定
FROM golang:latest

# Docker上の作業ディレクトリ設定
WORKDIR /go/src/work

# ホストのファイルをコンテナにコピー
ADD . /go/src/work

