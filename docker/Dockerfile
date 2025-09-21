# Multi-version Go Release Tour Dockerfile
# Supports Go versions 1.18 through 1.25 for accurate lesson execution

# Build stage - アプリケーションをビルド
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

# Multi-Go runtime stage - 複数のGoバージョンをインストール
FROM debian:bookworm-slim AS multi-go

# 必要なツールをインストール
RUN apt-get update && apt-get install -y \
    wget \
    tar \
    ca-certificates \
    curl \
    && rm -rf /var/lib/apt/lists/*

# 各Goバージョンを並列でダウンロード・インストール
ARG GO_VERSIONS="1.18.10 1.19.13 1.20.14 1.21.12 1.22.7 1.23.1 1.24.0 1.25.1"

# プラットフォーム検出とGoアーキテクチャ決定
RUN mkdir -p /opt/go && \
    ARCH=$(dpkg --print-architecture) && \
    if [ "$ARCH" = "amd64" ]; then \
        GO_ARCH="amd64"; \
    elif [ "$ARCH" = "arm64" ]; then \
        GO_ARCH="arm64"; \
    else \
        echo "Unsupported architecture: $ARCH" && exit 1; \
    fi && \
    echo "Detected architecture: $ARCH, using Go arch: $GO_ARCH" && \
    for version in $GO_VERSIONS; do \
        echo "Installing Go ${version} (${GO_ARCH})..." && \
        major_minor=$(echo $version | cut -d. -f1,2) && \
        wget -q -O- "https://go.dev/dl/go${version}.linux-${GO_ARCH}.tar.gz" | \
        tar -xz -C /tmp && \
        mv /tmp/go "/opt/go${major_minor}" && \
        echo "Go ${version} installed at /opt/go${major_minor}" \
    ; done

# バージョン確認とパス設定
RUN for dir in /opt/go*; do \
        version_path="$dir/bin/go" && \
        if [ -x "$version_path" ]; then \
            echo "Verifying: $version_path" && \
            $version_path version \
        ; fi \
    ; done

# デフォルトのGoを最新バージョンに設定
ENV PATH="/opt/go1.25/bin:$PATH"

# 最終ステージ - アプリケーション実行環境
FROM multi-go AS runtime

WORKDIR /app

# アプリケーションファイルをコピー
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static
COPY --from=builder /app/releases ./releases
COPY --from=builder /app/config ./config

# Go実行環境の確認
RUN echo "Available Go versions:" && \
    ls -la /opt/ && \
    echo "Default Go version:" && \
    go version

# ヘルスチェック設定
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/api/versions || exit 1

EXPOSE 8080

# 環境変数でマルチバージョン対応を示す
ENV MULTI_GO_SUPPORT=true
ENV GO_VERSIONS_AVAILABLE="1.18,1.19,1.20,1.21,1.22,1.23,1.24,1.25"

CMD ["./main"]