#!/bin/bash

# GitSentry Uninstall Script for Unix-like systems

set -e

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

# Remove binary
remove_binary() {
    local binary_path="$INSTALL_DIR/$BINARY_NAME"
    
    if [ -f "$binary_path" ]; then
        rm "$binary_path"
        print_success "Removed GitSentry binary from $binary_path"
    else
        print_warning "GitSentry binary not found at $binary_path"
    fi
}

# Remove PATH entry (optional)
remove_from_path() {
    local shell_rc=""
    
    # Detect shell and set appropriate RC file
    case "$SHELL" in
        */bash) shell_rc="$HOME/.bashrc" ;;
        */zsh)  shell_rc="$HOME/.zshrc" ;;
        */fish) shell_rc="$HOME/.config/fish/config.fish" ;;
        *)      shell_rc="$HOME/.profile" ;;
    esac
    
    if [ -f "$shell_rc" ] && grep -q "GitSentry" "$shell_rc"; then
        print_status "Found GitSentry PATH entry in $shell_rc"
        read -p "Remove PATH entry from $shell_rc? (y/N): " -n 1 -r
        echo
        
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            # Create backup
            cp "$shell_rc" "$shell_rc.backup"
            
            # Remove GitSentry lines
            sed -i '/# GitSentry/,+1d' "$shell_rc" 2>/dev/null || \
            sed -i.bak '/# GitSentry/,+1d' "$shell_rc"
            
            print_success "Removed PATH entry from $shell_rc"
            print_status "Backup created at $shell_rc.backup"
        fi
    fi
}

# Clean up .gitsentry directories (optional)
cleanup_projects() {
    print_status "Checking for .gitsentry directories in common project locations..."
    
    local found_dirs=()
    
    # Look for .gitsentry directories
    if command -v find >/dev/null 2>&1; then
        while IFS= read -r -d '' dir; do
            found_dirs+=("$dir")
        done < <(find "$HOME" -name ".gitsentry" -type d -print0 2>/dev/null | head -20)
    fi
    
    if [ ${#found_dirs[@]} -gt 0 ]; then
        echo "Found .gitsentry directories:"
        for dir in "${found_dirs[@]}"; do
            echo "  $dir"
        done
        
        read -p "Remove all .gitsentry project directories? (y/N): " -n 1 -r
        echo
        
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            for dir in "${found_dirs[@]}"; do
                rm -rf "$dir"
                print_success "Removed $dir"
            done
        fi
    else
        print_status "No .gitsentry directories found"
    fi
}

# Main uninstall process
main() {
    print_status "Starting GitSentry uninstallation..."
    
    # Remove binary
    remove_binary
    
    # Remove from PATH
    remove_from_path
    
    # Clean up project directories
    cleanup_projects
    
    print_success "GitSentry uninstallation completed!"
    print_status "Thank you for using GitSentry!"
}

# Run main function
main "$@"