BINARY_NAME=build/main

.PHONY: all build run clean docker-build docker-run docker-clean

all: run

deps:
	@echo "==> Installing dependencies..."
	go mod tidy

build: deps
	@echo "==> Building the application..."
	mkdir build
	go build -o $(BINARY_NAME) cmd/shortlink/main.go

run: build
	@echo "==> Running the application..."
	./$(BINARY_NAME)

clean:
	@echo "==> Cleaning up..."
	go clean
	rm -f $(BINARY_NAME)
	rm -rf build

init-db:
	@echo "==> Initializing database..."
	docker compose exec db psql -U user -d shortlink -f /scripts/create_table.sql

docker-build:
	@echo "==> Building Docker containers..."
	docker compose build

docker-up: docker-build
	@echo "==> Starting Docker containers..."
	docker compose up

docker-down:
	@echo "==> Stopping Docker containers..."
	docker compose down