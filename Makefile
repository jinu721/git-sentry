.PHONY: build clean install test run

# Build the application
build:
	go build -o bin/gitsentry cmd/gitsentry/main.go

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run tests
test:
	go test ./...

# Run the application
run:
	go run cmd/gitsentry/main.go

# Install to system (Unix-like systems)
install: build
	sudo cp bin/gitsentry /usr/local/bin/
	@echo "✅ GitSentry installed globally!"
	@echo "You can now use 'gitsentry' from any directory"

# Install to system (Windows - requires admin PowerShell)
install-windows: build-windows
	@echo "Installing GitSentry to C:\Windows\System32..."
	@echo "Note: Run this in Administrator PowerShell"
	copy bin\gitsentry.exe C:\Windows\System32\
	@echo "✅ GitSentry installed globally!"
	@echo "You can now use 'gitsentry' from any directory"

# Build for Windows
build-windows:
	go build -o bin/gitsentry.exe cmd/gitsentry/main.go

# Development build with race detection
dev:
	go build -race -o bin/gitsentry-dev cmd/gitsentry/main.go

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Show help
help:
	@echo "Available targets:"
	@echo "  build         - Build the application"
	@echo "  build-windows - Build for Windows"
	@echo "  clean         - Clean build artifacts"
	@echo "  deps          - Install dependencies"
	@echo "  test          - Run tests"
	@echo "  run           - Run the application"
	@echo "  install       - Install to system (Unix/Linux/macOS)"
	@echo "  install-windows - Install to system (Windows)"
	@echo "  dev           - Development build"
	@echo "  fmt           - Format code"
	@echo "  lint          - Lint code"
	@echo "  help          - Show this help"