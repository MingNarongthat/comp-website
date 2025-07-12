.PHONY: up down build rebuild clean

# Start all services
up:
	docker-compose up -d

# Stop all services
down:
	docker-compose down

# Build all services
build:
	docker-compose build

# Rebuild and restart all services
rebuild:
	docker-compose down
	docker-compose build
	docker-compose up -d

# Clean all docker volumes
clean:
	docker-compose down -v
	docker system prune -f

# Run backend development server
backend-dev:
	cd backend && go run cmd/api/main.go

# Run frontend development server
frontend-dev:
	cd frontend && npm run dev

# Install frontend dependencies
frontend-install:
	cd frontend && npm install

# Run database migrations
migrate:
	cd backend && go run cmd/migration/main.go

# Generate API documentation
docs:
	cd backend && swag init -g cmd/api/main.go

# Run backend tests
test:
	cd backend && go test ./...

# Format code
fmt:
	cd backend && go fmt ./...
	cd frontend && npm run lint:fix

# Check code style
lint:
	cd backend && golangci-lint run
	cd frontend && npm run lint
