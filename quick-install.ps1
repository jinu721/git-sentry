# GitSentry Quick Install Script for Windows
# Usage: iwr -useb https://raw.githubusercontent.com/jinu721/git-sentry/main/quick-install.ps1 | iex

Write-Host "Installing GitSentry..." -ForegroundColor Green

# Check if running as Administrator
if (-NOT ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) {
    Write-Host "Error: This script must be run as Administrator" -ForegroundColor Red
    Write-Host "Right-click PowerShell and select 'Run as Administrator'" -ForegroundColor Yellow
    exit 1
}

# Check if Go is installed
try {
    $goVersion = go version
    Write-Host "Found Go: $goVersion" -ForegroundColor Green
} catch {
    Write-Host "Error: Go is required but not installed." -ForegroundColor Red
    Write-Host "Please install Go from https://golang.org/doc/install" -ForegroundColor Yellow
    exit 1
}

# Create temporary directory
$tempDir = New-TemporaryFile | %{ rm $_; mkdir $_ }
Set-Location $tempDir

try {
    # Clone repository
    Write-Host "Downloading GitSentry..." -ForegroundColor Blue
    git clone https://github.com/jinu721/git-sentry.git
    Set-Location git-sentry

    # Build
    Write-Host "Building GitSentry..." -ForegroundColor Blue
    go build -o gitsentry.exe cmd/gitsentry/main.go

    # Install to System32
    $installDir = "C:\Windows\System32"
    Write-Host "Installing to $installDir..." -ForegroundColor Blue
    Copy-Item "gitsentry.exe" "$installDir\gitsentry.exe" -Force

    Write-Host "✅ GitSentry installed successfully!" -ForegroundColor Green
    Write-Host ""
    Write-Host "Usage:" -ForegroundColor Cyan
    Write-Host "  gitsentry init --template=team    # Initialize in your project" -ForegroundColor Cyan
    Write-Host "  gitsentry start --daemon          # Start monitoring" -ForegroundColor Cyan
    Write-Host "  gitsentry doctor                  # Run diagnostics" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "Run 'gitsentry --help' for more commands." -ForegroundColor Yellow

} catch {
    Write-Host "Installation failed: $_" -ForegroundColor Red
    exit 1
} finally {
    # Cleanup
    Set-Location C:\
    Remove-Item $tempDir -Recurse -Force -ErrorAction SilentlyContinue
}