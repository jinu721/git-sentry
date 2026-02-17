# GitSentry Windows Uninstall Script

$ErrorActionPreference = "Stop"

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

function Remove-Binary {
    $binaryPath = Join-Path $INSTALL_DIR $BINARY_NAME
    
    if (Test-Path $binaryPath) {
        Remove-Item $binaryPath -Force
        Write-Success "Removed GitSentry binary from $binaryPath"
    }
    else {
        Write-Warning "GitSentry binary not found at $binaryPath"
    }
}

function Remove-FromPath {
    $currentPath = [Environment]::GetEnvironmentVariable("PATH", "User")
    
    if ($currentPath -and ($currentPath -split ";" -contains $INSTALL_DIR)) {
        $choice = Read-Host "Remove $INSTALL_DIR from PATH? (y/N)"
        
        if ($choice -eq "y" -or $choice -eq "Y") {
            $newPath = ($currentPath -split ";" | Where-Object { $_ -ne $INSTALL_DIR }) -join ";"
            [Environment]::SetEnvironmentVariable("PATH", $newPath, "User")
            Write-Success "Removed $INSTALL_DIR from user PATH"
        }
    }
    else {
        Write-Status "GitSentry not found in user PATH"
    }
}

function Remove-ProjectDirectories {
    Write-Status "Checking for .gitsentry directories..."
    
    $gitsentryDirs = @()
    
    # Look for .gitsentry directories in common locations
    $searchPaths = @(
        "$env:USERPROFILE\Documents",
        "$env:USERPROFILE\Desktop",
        "$env:USERPROFILE\Projects"
    )
    
    foreach ($searchPath in $searchPaths) {
        if (Test-Path $searchPath) {
            $dirs = Get-ChildItem -Path $searchPath -Recurse -Directory -Name ".gitsentry" -ErrorAction SilentlyContinue
            foreach ($dir in $dirs) {
                $gitsentryDirs += Join-Path $searchPath $dir
            }
        }
    }
    
    if ($gitsentryDirs.Count -gt 0) {
        Write-Host "Found .gitsentry directories:"
        foreach ($dir in $gitsentryDirs) {
            Write-Host "  $dir"
        }
        
        $choice = Read-Host "Remove all .gitsentry project directories? (y/N)"
        
        if ($choice -eq "y" -or $choice -eq "Y") {
            foreach ($dir in $gitsentryDirs) {
                Remove-Item $dir -Recurse -Force -ErrorAction SilentlyContinue
                Write-Success "Removed $dir"
            }
        }
    }
    else {
        Write-Status "No .gitsentry directories found"
    }
}

function Main {
    Write-Status "Starting GitSentry uninstallation..."
    
    # Remove binary
    Remove-Binary
    
    # Remove from PATH
    Remove-FromPath
    
    # Clean up project directories
    Remove-ProjectDirectories
    
    Write-Success "GitSentry uninstallation completed!"
    Write-Status "Thank you for using GitSentry!"
}

# Run main function
Main