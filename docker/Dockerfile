FROM golang:1.25-alpine AS builder

WORKDIR /app

# go.modとgo.sumが存在する場合はコピー
COPY go.* ./

# go.modが存在しない場合はモジュールを初期化
RUN if [ ! -f go.mod ]; then go mod init go-release-tour; fi

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main ./app/cmd/server

# 最終ステージ
FROM golang:1.25-alpine

WORKDIR /app

# ヘルスチェック用にcurlをインストール
RUN apk add --no-cache curl

# バイナリと静的ファイルをコピー
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static
COPY --from=builder /app/releases ./releases

EXPOSE 8080

CMD ["./main"]