.PHONY: build clean test run help install-deps build-all release

# Build the application
build:
	go build -ldflags="-s -w" -o bin/gitsentry cmd/gitsentry/main.go

# Build for all platforms
build-all:
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o dist/gitsentry-linux-amd64 cmd/gitsentry/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o dist/gitsentry-linux-arm64 cmd/gitsentry/main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o dist/gitsentry-darwin-amd64 cmd/gitsentry/main.go
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o dist/gitsentry-darwin-arm64 cmd/gitsentry/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/gitsentry-windows-amd64.exe cmd/gitsentry/main.go
	GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -o dist/gitsentry-windows-arm64.exe cmd/gitsentry/main.go
	GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o dist/gitsentry-freebsd-amd64 cmd/gitsentry/main.go

# Clean build artifacts
clean:
	rm -rf bin/ dist/

# Install dependencies
install-deps:
	go mod tidy
	go mod download

# Run tests
test:
	go test ./internal/...

# Run the application
run:
	go run cmd/gitsentry/main.go

# Install to user directory (cross-platform)
install:
	@echo "Use the installation script instead:"
	@echo "  curl -sSL https://raw.githubusercontent.com/jinu721/git-sentry/main/scripts/install.sh | bash"

# Development build with race detection
dev:
	go build -race -o bin/gitsentry-dev cmd/gitsentry/main.go

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Create a release (requires git tag)
release: build-all
	@echo "Built binaries for release in dist/"
	@echo "Create a git tag and push to trigger GitHub Actions release"

# Show help
help:
	@echo "Available targets:"
	@echo "  build         - Build the application for current platform"
	@echo "  build-all     - Build for all supported platforms"
	@echo "  clean         - Clean build artifacts"
	@echo "  install-deps  - Install Go dependencies"
	@echo "  test          - Run tests"
	@echo "  run           - Run the application"
	@echo "  install       - Show installation instructions"
	@echo "  dev           - Development build with race detection"
	@echo "  fmt           - Format code"
	@echo "  lint          - Lint code"
	@echo "  release       - Build all platforms for release"
	@echo "  help          - Show this help"