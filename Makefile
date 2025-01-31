APP_NAME=go-browser
BUILD_DIR=bin
GO_FILES=$(shell find . -name '*.go' -type f)

.PHONY: all build run clean test

all: build

build:
	@echo "==> Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) ./main.go

run: build
	@echo "==> Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

clean:
	@echo "==> Cleaning up..."
	@rm -rf $(BUILD_DIR)

test:
	@echo "==> Running tests..."
	@go test -v ./tests/...

fmt:
	@echo "==> Formatting code..."
	@go fmt ./tests

lint:
	@echo "==> Running golangci-lint..."
	@golangci-lint run

vet:
	@echo "==> Running go vet..."
	@go vet ./...

help:
	@echo "Available commands:"
	@echo "  make build   - Build the application"
	@echo "  make run     - Run the application"
	@echo "  make clean   - Remove built files"
	@echo "  make test    - Run tests"
	@echo "  make fmt     - Format the code"
	@echo "  make lint    - Run golangci-lint"
	@echo "  make vet     - Run go vet"
