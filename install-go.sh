#!/bin/bash

echo "🚀 GitSentry Installation via Go Install"
echo "========================================"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go first."
    echo "Visit: https://golang.org/doc/install"
    exit 1
fi

echo "✅ Go found: $(go version)"

# Install GitSentry using go install
echo "📦 Installing GitSentry globally..."
go install ./cmd/gitsentry

# Check if GOPATH/bin is in PATH
if ! command -v gitsentry &> /dev/null; then
    echo "⚠️  GitSentry installed but not in PATH"
    echo ""
    echo "Add Go's bin directory to your PATH:"
    echo "  export PATH=\$PATH:\$(go env GOPATH)/bin"
    echo ""
    echo "Or add this line to your ~/.bashrc or ~/.zshrc:"
    echo "  export PATH=\$PATH:\$(go env GOPATH)/bin"
    echo ""
    echo "Then restart your terminal or run: source ~/.bashrc"
else
    echo "✅ GitSentry installed successfully!"
    echo ""
    echo "🎉 You can now use 'gitsentry' from any directory:"
    echo "   gitsentry init    # Initialize in any project"
    echo "   gitsentry start   # Start monitoring"
    echo "   gitsentry status  # Check status"
    echo ""
    echo "📚 Run 'gitsentry --help' for more commands"
fi