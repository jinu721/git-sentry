#!/bin/bash

echo "🗑️  GitSentry Uninstall Script"
echo "============================="

# Check common installation locations
LOCATIONS=(
    "/usr/local/bin/gitsentry"
    "$(go env GOPATH)/bin/gitsentry"
    "$HOME/.local/bin/gitsentry"
)

FOUND=false

for location in "${LOCATIONS[@]}"; do
    if [ -f "$location" ]; then
        echo "📍 Found GitSentry at: $location"
        
        if [[ "$location" == "/usr/local/bin/gitsentry" ]]; then
            echo "🔐 Removing system installation (requires sudo)..."
            sudo rm "$location"
        else
            echo "🗑️  Removing user installation..."
            rm "$location"
        fi
        
        echo "✅ Removed: $location"
        FOUND=true
    fi
done

if [ "$FOUND" = false ]; then
    echo "❌ GitSentry not found in common locations"
    echo "Check your PATH and remove manually if needed"
    exit 1
fi

echo ""
echo "✅ GitSentry uninstalled successfully!"
echo "Note: This doesn't remove .gitsentry/ folders from your projects"