# Go Release Tour - Makefile

.PHONY: help dev prod build clean install air docker-dev docker-prod docker-multi stop

# デフォルトターゲット
help:
	@echo "Go Release Tour - Available commands:"
	@echo ""
	@echo "Development:"
	@echo "  make dev         - Start development server with Air (hot reload)"
	@echo "  make air         - Install Air and start development server"
	@echo "  make docker-dev  - Start development environment with Docker"
	@echo ""
	@echo "Production:"
	@echo "  make prod        - Start production server"
	@echo "  make docker-prod - Start production environment with Docker"
	@echo "  make docker-multi- Start multi-version environment with proxy"
	@echo "  make build       - Build binary for production"
	@echo ""
	@echo "Utilities:"
	@echo "  make install     - Install dependencies and tools"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make stop        - Stop all Docker containers"
	@echo "  make help        - Show this help message"

# 開発環境
dev: air

air:
	@echo "Starting development server with Air (hot reload)..."
	@command -v air >/dev/null 2>&1 || (echo "Installing Air..." && go install github.com/air-verse/air@latest)
	cd .. && air -c build/.air.toml

docker-dev:
	@echo "Starting development environment with Docker..."
	cd .. && docker-compose -f docker-compose.dev.yml up

# 本番環境
prod: build
	@echo "Starting production server..."
	./go-release-tour

docker-prod:
	@echo "Starting production environment with Docker..."
	cd .. && docker-compose up -d

docker-multi:
	@echo "Starting multi-version environment with proxy..."
	docker-compose -f docker-compose.multiversion.yml up -d

build:
	@echo "Building binary..."
	cd .. && go build -o go-release-tour ./app/cmd/server

# ユーティリティ
install:
	@echo "Installing dependencies and tools..."
	go mod download
	go install github.com/air-verse/air@latest
	@echo "Installation complete!"

clean:
	@echo "Cleaning build artifacts..."
	rm -rf tmp/
	rm -f go-release-tour
	rm -f build-errors.log
	docker system prune -f
	@echo "Clean complete!"

stop:
	@echo "Stopping all containers..."
	docker-compose down
	docker-compose -f docker-compose.dev.yml down
	docker-compose -f docker-compose.multiversion.yml down
	@echo "All containers stopped!"

# クイックコマンド
start: dev
restart: stop dev
multi: docker-multi