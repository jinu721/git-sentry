#!/bin/bash

set -e

echo "🚀 GitSentry Global Installation Script"
echo "======================================"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go first."
    echo "Visit: https://golang.org/doc/install"
    exit 1
fi

echo "✅ Go found: $(go version)"

# Build GitSentry
echo "🔨 Building GitSentry..."
go build -o gitsentry cmd/gitsentry/main.go

# Determine installation directory
if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS
    INSTALL_DIR="/usr/local/bin"
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    # Linux
    INSTALL_DIR="/usr/local/bin"
else
    echo "❌ Unsupported operating system: $OSTYPE"
    echo "Please install manually by copying 'gitsentry' to your PATH"
    exit 1
fi

# Check if installation directory exists
if [ ! -d "$INSTALL_DIR" ]; then
    echo "❌ Installation directory $INSTALL_DIR does not exist"
    exit 1
fi

# Install GitSentry
echo "📦 Installing GitSentry to $INSTALL_DIR..."
if [ "$EUID" -eq 0 ]; then
    # Running as root
    cp gitsentry "$INSTALL_DIR/"
    chmod +x "$INSTALL_DIR/gitsentry"
else
    # Need sudo
    sudo cp gitsentry "$INSTALL_DIR/"
    sudo chmod +x "$INSTALL_DIR/gitsentry"
fi

# Clean up
rm gitsentry

echo "✅ GitSentry installed successfully!"
echo ""
echo "🎉 You can now use 'gitsentry' from any directory:"
echo "   gitsentry init    # Initialize in any project"
echo "   gitsentry start   # Start monitoring"
echo "   gitsentry status  # Check status"
echo ""
echo "📚 Run 'gitsentry --help' for more commands"