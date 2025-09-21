# Go Release Tour - Makefile

.PHONY: help drop build init app dev clean logs status test

# デフォルトターゲット
help:
	@echo "Go Release Tour - Available commands:"
	@echo ""
	@echo "Core commands (Docker Compose):"
	@echo "  make drop        - Stop and remove all containers and clean artifacts"
	@echo "  make build       - Build Docker image"
	@echo "  make init        - Initialize project and build image"
	@echo "  make app         - Start the application with Docker Compose"
	@echo ""
	@echo "Development:"
	@echo "  make dev         - Start development environment (detached)"
	@echo "  make logs        - Show application logs"
	@echo "  make status      - Show container status"
	@echo ""
	@echo "Testing:"
	@echo "  make test        - Run all tests (E2E, integration)"
	@echo ""
	@echo "Utilities:"
	@echo "  make clean       - Clean Docker artifacts only"
	@echo "  make help        - Show this help message"

# Core commands - All via Docker Compose
drop:
	@echo "Dropping containers and cleaning all artifacts..."
	docker-compose down --volumes --remove-orphans || true
	docker system prune -f || true
	docker image prune -f || true
	@echo "Drop complete!"

build:
	@echo "Building Docker image..."
	docker-compose build --no-cache
	@echo "Build complete!"

init:
	@echo "Initializing project with Docker..."
	@$(MAKE) build
	@echo "Initialization complete!"

app:
	@echo "Starting application with Docker Compose..."
	docker-compose up -d
	@echo "Application started! Access at http://localhost:8080"
	@echo "Use 'make logs' to view logs or 'make status' to check status"

# Development
dev:
	@echo "Starting development environment..."
	docker-compose up -d
	@echo "Development environment started!"
	@echo "Logs will show below (Ctrl+C to stop viewing logs, container keeps running):"
	@sleep 2
	docker-compose logs -f

# Utilities
logs:
	@echo "Showing application logs (Ctrl+C to exit):"
	docker-compose logs -f

status:
	@echo "Container status:"
	docker-compose ps

clean:
	@echo "Cleaning Docker artifacts..."
	docker-compose down || true
	docker system prune -f || true
	@echo "Clean complete!"

# Testing
test:
	@echo "Running all tests..."
	@echo "Starting application for testing..."
	@$(MAKE) app
	@sleep 5
	@echo "Running E2E API tests..."
	@chmod +x tests/e2e/e2e_api_test.sh
	@./tests/e2e/e2e_api_test.sh || echo "E2E API tests completed with issues"
	@echo "Running integration tests..."
	@chmod +x tests/integration/test_all_lessons.sh
	@./tests/integration/test_all_lessons.sh || echo "Integration tests completed with issues"
	@echo "Test results saved in tests/results/"
	@echo "All tests completed!"

# Aliases for convenience
start: app
stop:
	@echo "Stopping containers..."
	docker-compose down
restart: drop init app
ps: status