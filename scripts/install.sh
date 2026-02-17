#!/bin/bash

# GitSentry Universal Installation Script
# Works on Linux, macOS, FreeBSD, WSL
# No admin rights required - installs to user directory

set -e

REPO="jinu721/git-sentry"
INSTALL_DIR="$HOME/.local/bin"
BINARY_NAME="gitsentry"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Detect OS and architecture
detect_platform() {
    local os arch
    
    case "$(uname -s)" in
        Linux*)     os="linux" ;;
        Darwin*)    os="darwin" ;;
        FreeBSD*)   os="freebsd" ;;
        CYGWIN*|MINGW*|MSYS*) os="windows" ;;
        *)          print_error "Unsupported OS: $(uname -s)"; exit 1 ;;
    esac
    
    case "$(uname -m)" in
        x86_64|amd64)   arch="amd64" ;;
        arm64|aarch64)  arch="arm64" ;;
        *)              print_error "Unsupported architecture: $(uname -m)"; exit 1 ;;
    esac
    
    echo "${os}-${arch}"
}

# Get latest release version
get_latest_version() {
    local version
    version=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    
    if [ -z "$version" ]; then
        print_error "Failed to get latest version"
        exit 1
    fi
    
    echo "$version"
}

# Download and install binary
install_binary() {
    local platform="$1"
    local version="$2"
    local suffix=""
    
    if [[ "$platform" == *"windows"* ]]; then
        suffix=".exe"
    fi
    
    local binary_name="${BINARY_NAME}-${platform}${suffix}"
    local download_url="https://github.com/${REPO}/releases/download/${version}/${binary_name}"
    
    print_status "Downloading GitSentry ${version} for ${platform}..."
    
    # Create install directory
    mkdir -p "$INSTALL_DIR"
    
    # Download binary
    if command -v curl >/dev/null 2>&1; then
        curl -L "$download_url" -o "$INSTALL_DIR/$BINARY_NAME$suffix"
    elif command -v wget >/dev/null 2>&1; then
        wget "$download_url" -O "$INSTALL_DIR/$BINARY_NAME$suffix"
    else
        print_error "Neither curl nor wget found. Please install one of them."
        exit 1
    fi
    
    # Make executable
    chmod +x "$INSTALL_DIR/$BINARY_NAME$suffix"
    
    print_success "GitSentry installed to $INSTALL_DIR/$BINARY_NAME$suffix"
}

# Add to PATH if needed
setup_path() {
    local shell_rc=""
    
    # Detect shell and set appropriate RC file
    case "$SHELL" in
        */bash) shell_rc="$HOME/.bashrc" ;;
        */zsh)  shell_rc="$HOME/.zshrc" ;;
        */fish) shell_rc="$HOME/.config/fish/config.fish" ;;
        *)      shell_rc="$HOME/.profile" ;;
    esac
    
    # Check if already in PATH
    if echo "$PATH" | grep -q "$INSTALL_DIR"; then
        print_status "Install directory already in PATH"
        return
    fi
    
    # Add to PATH in shell RC file
    if [ -f "$shell_rc" ]; then
        if ! grep -q "$INSTALL_DIR" "$shell_rc"; then
            echo "" >> "$shell_rc"
            echo "# GitSentry" >> "$shell_rc"
            echo "export PATH=\"\$PATH:$INSTALL_DIR\"" >> "$shell_rc"
            print_success "Added $INSTALL_DIR to PATH in $shell_rc"
            print_warning "Please restart your terminal or run: source $shell_rc"
        fi
    else
        print_warning "Could not find shell RC file. Please add $INSTALL_DIR to your PATH manually."
    fi
}

# Verify installation
verify_installation() {
    if [ -x "$INSTALL_DIR/$BINARY_NAME" ]; then
        print_success "Installation verified!"
        print_status "GitSentry version: $($INSTALL_DIR/$BINARY_NAME --version 2>/dev/null || echo 'installed')"
    else
        print_error "Installation verification failed"
        exit 1
    fi
}

# Main installation process
main() {
    print_status "Starting GitSentry installation..."
    
    # Check dependencies
    if ! command -v curl >/dev/null 2>&1 && ! command -v wget >/dev/null 2>&1; then
        print_error "curl or wget is required for installation"
        exit 1
    fi
    
    # Detect platform
    local platform
    platform=$(detect_platform)
    print_status "Detected platform: $platform"
    
    # Get latest version
    local version
    version=$(get_latest_version)
    print_status "Latest version: $version"
    
    # Install binary
    install_binary "$platform" "$version"
    
    # Setup PATH
    setup_path
    
    # Verify installation
    verify_installation
    
    print_success "GitSentry installation completed!"
    echo ""
    print_status "Quick start:"
    echo "  cd /path/to/your/git/project"
    echo "  gitsentry init --template=team"
    echo "  gitsentry start --daemon"
    echo ""
    print_status "Run 'gitsentry --help' for more commands"
}

# Run main function
main "$@"