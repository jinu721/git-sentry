# GitSentry Windows Uninstall Script

Write-Host "🗑️  GitSentry Uninstall Script" -ForegroundColor Red
Write-Host "=============================" -ForegroundColor Red

# Check common installation locations
$locations = @(
    "C:\Windows\System32\gitsentry.exe",
    "$env:GOPATH\bin\gitsentry.exe",
    "$env:USERPROFILE\.local\bin\gitsentry.exe"
)

$found = $false

foreach ($location in $locations) {
    if (Test-Path $location) {
        Write-Host "📍 Found GitSentry at: $location" -ForegroundColor Yellow
        
        try {
            if ($location -like "C:\Windows\System32\*") {
                Write-Host "🔐 Removing system installation (requires admin)..." -ForegroundColor Blue
                Remove-Item $location -Force
            } else {
                Write-Host "🗑️  Removing user installation..." -ForegroundColor Blue
                Remove-Item $location -Force
            }
            
            Write-Host "✅ Removed: $location" -ForegroundColor Green
            $found = $true
        } catch {
            Write-Host "❌ Failed to remove: $location" -ForegroundColor Red
            Write-Host "Error: $_" -ForegroundColor Red
        }
    }
}

if (-not $found) {
    Write-Host "❌ GitSentry not found in common locations" -ForegroundColor Red
    Write-Host "Check your PATH and remove manually if needed" -ForegroundColor Yellow
    exit 1
}

Write-Host ""
Write-Host "✅ GitSentry uninstalled successfully!" -ForegroundColor Green
Write-Host "Note: This doesn't remove .gitsentry/ folders from your projects" -ForegroundColor Yellow