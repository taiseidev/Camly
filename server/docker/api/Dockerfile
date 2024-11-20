# ベースイメージ
FROM golang:1.21-alpine

# 必要なツールをインストール
RUN apk update && apk add --no-cache git

# 作業ディレクトリを設定
WORKDIR /app/cmd/api

# go.mod と go.sum をコンテナにコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY ./cmd/api/ .

# アプリケーションで使用するポートを公開
EXPOSE 8080

# Airをインストール
RUN go install github.com/cosmtrek/air@v1.40.4

# Airを使ってAPIサーバーを実行
CMD ["air", "-c", "../.air.toml"]
