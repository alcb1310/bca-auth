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
