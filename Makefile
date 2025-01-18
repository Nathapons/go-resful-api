# Simple Makefile for a Go project

# Build the application
all: build

# Migration Variable
MIGRATIONS_PATH=migrations
TYPE ?= sql

# Install project
install:
	@echo "Installing Goose..."

	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@if ! command -v migrate &> /dev/null; then \
		echo "Goose installation is failed!!!"; \
	else \
		echo "Goose installation is complete"; \
	fi

	@echo "Installing dependencies..."
	@go mod download
	@if ! command -v migrate &> /dev/null; then \
		echo "Project installation is failed!!!"; \
	else \
		echo "Project installation is complete"; \
	fi

build:
	@echo "Building..."
	@go build -o bin/api/main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		docker-compose up; \
		echo "Up docker-compose.yml complete"; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		docker-compose down; \
		echo "Down docker-compose.local.yml complete"; \
	fi

# Target to create a new Goose migration
create-migration:
	@if [ -z "$(NAME)" ]; then \
		echo "NAME is not set. Usage: make create-migration NAME=<migration_name> [TYPE=go|sql]"; \
		exit 1; \
	fi

	@if [ "$(TYPE)" != "sql" ] && [ "$(TYPE)" != "go" ]; then \
		echo "Invalid TYPE specified. Choose from 'sql' or 'go'"; \
		exit 1; \
	fi

	@mkdir -p $(MIGRATIONS_PATH)
	@goose -dir ./$(MIGRATIONS_PATH) create $(NAME) $(TYPE)

migrate-up:
	@if ! command -v goose &> /dev/null; then \
        echo "Goose is not installed. Please run 'make install' first."; \
        exit 1; \
    fi
	goose -dir ${MIGRATIONS_PATH} postgres "postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" up

migrate-down:
	@if ! command -v goose &> /dev/null; then \
        echo "Goose is not installed. Please run 'make install' first."; \
        exit 1; \
    fi
	goose -dir ${MIGRATIONS_PATH} postgres "postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" down