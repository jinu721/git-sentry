#!/bin/bash

# GitSentry Quick Install Script
# Usage: curl -sSL https://raw.githubusercontent.com/jinu721/git-sentry/main/quick-install.sh | bash

set -e

echo "Installing GitSentry..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is required but not installed."
    echo "Please install Go from https://golang.org/doc/install"
    exit 1
fi

# Create temporary directory
TEMP_DIR=$(mktemp -d)
cd "$TEMP_DIR"

# Clone repository
echo "Downloading GitSentry..."
git clone https://github.com/jinu721/git-sentry.git
cd git-sentry

# Build and install
echo "Building GitSentry..."
go build -o gitsentry cmd/gitsentry/main.go

# Install to system
if [[ "$OSTYPE" == "darwin"* ]] || [[ "$OSTYPE" == "linux-gnu"* ]]; then
    INSTALL_DIR="/usr/local/bin"
    echo "Installing to $INSTALL_DIR..."
    
    if [ "$EUID" -eq 0 ]; then
        cp gitsentry "$INSTALL_DIR/"
        chmod +x "$INSTALL_DIR/gitsentry"
    else
        sudo cp gitsentry "$INSTALL_DIR/"
        sudo chmod +x "$INSTALL_DIR/gitsentry"
    fi
else
    echo "Unsupported OS. Please install manually."
    exit 1
fi

# Cleanup
cd /
rm -rf "$TEMP_DIR"

echo "✅ GitSentry installed successfully!"
echo ""
echo "Usage:"
echo "  gitsentry init --template=team    # Initialize in your project"
echo "  gitsentry start --daemon          # Start monitoring"
echo "  gitsentry doctor                  # Run diagnostics"
echo ""
echo "Run 'gitsentry --help' for more commands."