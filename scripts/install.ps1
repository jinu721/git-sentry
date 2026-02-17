# GitSentry Windows Installation Script
# No admin rights required - installs to user directory

param(
    [string]$Version = "latest"
)

$ErrorActionPreference = "Stop"

$REPO = "jinu721/git-sentry"
$INSTALL_DIR = "$env:USERPROFILE\.local\bin"
$BINARY_NAME = "gitsentry.exe"

function Write-Status {
    param([string]$Message)
    Write-Host "[INFO] $Message" -ForegroundColor Blue
}

function Write-Success {
    param([string]$Message)
    Write-Host "[SUCCESS] $Message" -ForegroundColor Green
}

function Write-Warning {
    param([string]$Message)
    Write-Host "[WARNING] $Message" -ForegroundColor Yellow
}

function Write-Error {
    param([string]$Message)
    Write-Host "[ERROR] $Message" -ForegroundColor Red
}

function Get-Platform {
    $arch = if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }
    
    # Check for ARM64
    if ($env:PROCESSOR_ARCHITECTURE -eq "ARM64") {
        $arch = "arm64"
    }
    
    return "windows-$arch"
}

function Get-LatestVersion {
    try {
        $response = Invoke-RestMethod -Uri "https://api.github.com/repos/$REPO/releases/latest"
        return $response.tag_name
    }
    catch {
        Write-Error "Failed to get latest version: $_"
        exit 1
    }
}

function Install-Binary {
    param(
        [string]$Platform,
        [string]$Version
    )
    
    $binaryName = "gitsentry-$Platform.exe"
    $downloadUrl = "https://github.com/$REPO/releases/download/$Version/$binaryName"
    
    Write-Status "Downloading GitSentry $Version for $Platform..."
    
    # Create install directory
    if (!(Test-Path $INSTALL_DIR)) {
        New-Item -ItemType Directory -Path $INSTALL_DIR -Force | Out-Null
    }
    
    $targetPath = Join-Path $INSTALL_DIR $BINARY_NAME
    
    try {
        Invoke-WebRequest -Uri $downloadUrl -OutFile $targetPath
        Write-Success "GitSentry installed to $targetPath"
    }
    catch {
        Write-Error "Failed to download binary: $_"
        exit 1
    }
}

function Setup-Path {
    # Get current user PATH
    $currentPath = [Environment]::GetEnvironmentVariable("PATH", "User")
    
    # Check if already in PATH
    if ($currentPath -split ";" -contains $INSTALL_DIR) {
        Write-Status "Install directory already in PATH"
        return
    }
    
    # Add to user PATH
    $newPath = if ($currentPath) { "$currentPath;$INSTALL_DIR" } else { $INSTALL_DIR }
    
    try {
        [Environment]::SetEnvironmentVariable("PATH", $newPath, "User")
        Write-Success "Added $INSTALL_DIR to user PATH"
        Write-Warning "Please restart your terminal or PowerShell session"
        
        # Update current session PATH
        $env:PATH = "$env:PATH;$INSTALL_DIR"
    }
    catch {
        Write-Warning "Could not update PATH automatically. Please add $INSTALL_DIR to your PATH manually."
    }
}

function Test-Installation {
    $targetPath = Join-Path $INSTALL_DIR $BINARY_NAME
    
    if (Test-Path $targetPath) {
        Write-Success "Installation verified!"
        try {
            $versionOutput = & $targetPath --version 2>$null
            Write-Status "GitSentry version: $versionOutput"
        }
        catch {
            Write-Status "GitSentry installed successfully"
        }
    }
    else {
        Write-Error "Installation verification failed"
        exit 1
    }
}

function Main {
    Write-Status "Starting GitSentry installation..."
    
    # Detect platform
    $platform = Get-Platform
    Write-Status "Detected platform: $platform"
    
    # Get version
    $version = if ($Version -eq "latest") { Get-LatestVersion } else { $Version }
    Write-Status "Installing version: $version"
    
    # Install binary
    Install-Binary -Platform $platform -Version $version
    
    # Setup PATH
    Setup-Path
    
    # Verify installation
    Test-Installation
    
    Write-Success "GitSentry installation completed!"
    Write-Host ""
    Write-Status "Quick start:"
    Write-Host "  cd C:\path\to\your\git\project"
    Write-Host "  gitsentry init --template=team"
    Write-Host "  gitsentry start --daemon"
    Write-Host ""
    Write-Status "Run 'gitsentry --help' for more commands"
}

# Run main function
Main