APP_NAME := evm-disassembler
PKG := ./...
CMD_DIR := ./cmd/server
BIN_DIR := bin

.PHONY: help build run test clean docker lint

## Show help
help:
	@echo "Available make commands:"
	@echo "  make build      - Build the Go binary"
	@echo "  make run        - Run the application locally"
	@echo "  make test       - Run unit tests"
	@echo "  make lint       - Run linters/checkers (if installed)"
	@echo "  make docker     - Build a Docker image"
	@echo "  make clean      - Remove temporary files/binaries"

## Build the Go binary
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)

## Run the application (development)
run:
	@echo "Running $(APP_NAME)..."
	@go run $(CMD_DIR)/main.go

## Run tests
test:
	@echo "Running tests..."
	@go test $(PKG) -v

## Run linter (requires golangci-lint or similar installed)
lint:
	@echo "Running linters..."
	@golangci-lint run

## Build the Docker image
docker:
	@echo "Building Docker image for $(APP_NAME)..."
	@docker build -t $(APP_NAME):latest .

## Clean up build files
clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)/$(APP_NAME)
