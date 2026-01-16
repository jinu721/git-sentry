# GitSentry Windows Installation Script
# Run this in Administrator PowerShell

Write-Host "🚀 GitSentry Global Installation Script" -ForegroundColor Green
Write-Host "======================================" -ForegroundColor Green

# Check if running as Administrator
if (-NOT ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) {
    Write-Host "❌ This script must be run as Administrator" -ForegroundColor Red
    Write-Host "Right-click PowerShell and select 'Run as Administrator'" -ForegroundColor Yellow
    exit 1
}

# Check if Go is installed
try {
    $goVersion = go version
    Write-Host "✅ Go found: $goVersion" -ForegroundColor Green
} catch {
    Write-Host "❌ Go is not installed. Please install Go first." -ForegroundColor Red
    Write-Host "Visit: https://golang.org/doc/install" -ForegroundColor Yellow
    exit 1
}

# Build GitSentry
Write-Host "🔨 Building GitSentry..." -ForegroundColor Blue
go build -o gitsentry.exe cmd/gitsentry/main.go

if (-not (Test-Path "gitsentry.exe")) {
    Write-Host "❌ Build failed" -ForegroundColor Red
    exit 1
}

# Install to System32 (in PATH by default)
$installDir = "C:\Windows\System32"
Write-Host "📦 Installing GitSentry to $installDir..." -ForegroundColor Blue

try {
    Copy-Item "gitsentry.exe" "$installDir\gitsentry.exe" -Force
    Write-Host "✅ GitSentry installed successfully!" -ForegroundColor Green
} catch {
    Write-Host "❌ Installation failed: $_" -ForegroundColor Red
    exit 1
}

# Clean up
Remove-Item "gitsentry.exe"

Write-Host ""
Write-Host "🎉 You can now use 'gitsentry' from any directory:" -ForegroundColor Green
Write-Host "   gitsentry init    # Initialize in any project" -ForegroundColor Cyan
Write-Host "   gitsentry start   # Start monitoring" -ForegroundColor Cyan
Write-Host "   gitsentry status  # Check status" -ForegroundColor Cyan
Write-Host ""
Write-Host "📚 Run 'gitsentry --help' for more commands" -ForegroundColor Yellow

# Test installation
Write-Host ""
Write-Host "🧪 Testing installation..." -ForegroundColor Blue
try {
    & gitsentry --help | Out-Null
    Write-Host "✅ Installation test passed!" -ForegroundColor Green
} catch {
    Write-Host "⚠️  Installation test failed. You may need to restart your terminal." -ForegroundColor Yellow
}