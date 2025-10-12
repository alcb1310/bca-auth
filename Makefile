# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o out cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Live Reload
watch:
	@air

clean:
	@echo "Cleaning..."
	@go clean -testcache
	@echo "Cache cleaned..."

# Test the application
unit-test: clean
	@echo "Unit Tests..."
	@go test `go list ./... | grep -v ./cmd/api | grep -v ./internal/database | grep -v ./mocks | grep -v ./tests | grep -v ./internal/views`

test: unit-test
