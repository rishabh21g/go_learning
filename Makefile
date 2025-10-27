# Go Learning Project Makefile
# This Makefile provides convenient commands for building, running, and testing the Go learning project

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet

# Binary name
BINARY_NAME=go_learning
BINARY_UNIX=$(BINARY_NAME)_unix

# Default target
.PHONY: all
all: clean build

# Build the application
.PHONY: build
build:
	@echo "🔨 Building Go Learning application..."
	$(GOBUILD) -o $(BINARY_NAME) -v

# Run the application
.PHONY: run
run:
	@echo "🚀 Running Go Learning application..."
	$(GOCMD) run main.go

# Run with basic examples
.PHONY: run-basics
run-basics:
	@echo "🏗️ Running basic syntax examples..."
	$(GOCMD) run basics/*.go

# Run function examples
.PHONY: run-functions
run-functions:
	@echo "⚙️ Running function examples..."
	$(GOCMD) run functions/*.go

# Run struct examples
.PHONY: run-structs
run-structs:
	@echo "🏛️ Running struct examples..."
	$(GOCMD) run structs/*.go

# Run concurrency examples
.PHONY: run-concurrency
run-concurrency:
	@echo "🚦 Running concurrency examples..."
	$(GOCMD) run concurrency/*.go

# Run backend examples
.PHONY: run-backend
run-backend:
	@echo "🌐 Running backend examples..."
	$(GOCMD) run backend/*.go

# Clean build artifacts
.PHONY: clean
clean:
	@echo "🧹 Cleaning build artifacts..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Format Go code
.PHONY: fmt
fmt:
	@echo "📝 Formatting Go code..."
	$(GOFMT) ./...

# Vet Go code
.PHONY: vet
vet:
	@echo "🔍 Vetting Go code..."
	$(GOVET) ./...

# Download dependencies
.PHONY: deps
deps:
	@echo "📦 Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) verify

# Tidy dependencies
.PHONY: tidy
tidy:
	@echo "🧹 Tidying Go modules..."
	$(GOMOD) tidy

# Check code quality
.PHONY: check
check: fmt vet
	@echo "✅ Code quality check completed"

# Build for Linux
.PHONY: build-linux
build-linux:
	@echo "🐧 Building for Linux..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

# Show help
.PHONY: help
help:
	@echo "🚀 Go Learning Project - Available Commands:"
	@echo ""
	@echo "📋 Basic Commands:"
	@echo "  make build          - Build the application"
	@echo "  make run            - Run the interactive application"
	@echo "  make clean          - Clean build artifacts"
	@echo ""
	@echo "🎓 Learning Commands:"
	@echo "  make run-basics     - Run basic syntax examples"
	@echo "  make run-functions  - Run function examples"
	@echo "  make run-structs    - Run struct and interface examples"
	@echo "  make run-concurrency - Run concurrency examples"
	@echo "  make run-backend    - Run backend server examples"
	@echo ""
	@echo "🔧 Development Commands:"
	@echo "  make fmt            - Format Go code"
	@echo "  make vet            - Vet Go code for issues"
	@echo "  make check          - Run fmt and vet"
	@echo "  make tidy           - Tidy Go modules"
	@echo "  make deps           - Download dependencies"
	@echo ""
	@echo "🌍 Build Commands:"
	@echo "  make build-linux    - Build for Linux"
	@echo ""
	@echo "❓ Help:"
	@echo "  make help           - Show this help message"

# Quick demo target
.PHONY: demo
demo:
	@echo "🎬 Running quick demo of all modules..."
	@echo "This will show a brief example from each learning module:"
	@echo ""
	@echo "1. Basic Syntax:"
	@$(GOCMD) run -c 'package main; import "fmt"; func main() { name := "Go"; fmt.Printf("Hello, %s!\n", name) }'
	@echo ""
	@echo "2. Full Interactive Demo:"
	@echo "   Run 'make run' to start the interactive learning application"
	@echo ""
	@echo "3. Individual Modules:"
	@echo "   Use 'make run-basics', 'make run-functions', etc."