param (
    [string]$RepositoryRoot = "./../..",
    [string]$GoVersion = "go1.22.4",
    [string]$Version = "1.0.0"
)

# Path to the go.work file
$goWorkFile = Join-Path $RepositoryRoot "go.work"

# Ensure the go.work file exists
if (-Not (Test-Path $goWorkFile)) {
    Write-Error "go.work file not found at $goWorkFile"
    exit 1
}

# Extract module paths from the go.work file
$modulePaths = Get-Content $goWorkFile |
    ForEach-Object { $_.Trim() } |
    Where-Object { $_ -match '^\.\/' -or $_ -match '^use ' } |
    ForEach-Object { $_ -replace '^use ', '' -replace '[()]', '' } |
    Where-Object { $_ -match '^\.\/' }

if ($modulePaths.Count -eq 0) {
    Write-Error "No modules found in the 'use' section of go.work"
    exit 1
}

Write-Host "Modules in solution: $modulePaths"

# Iterate over each module and build it
foreach ($module in $modulePaths) {
    $modulePath = Join-Path $RepositoryRoot ($module.Trim('"').Trim('./'))
    $publishDirectory = "bin/Release/$GoVersion/"

    if (-Not (Test-Path $modulePath)) {
        Write-Warning "Module path $modulePath does not exist. Skipping..."
        continue
    }

    Write-Host "Building module at $modulePath"
    Push-Location $modulePath

    # Run `go mod tidy` to ensure dependencies are resolved
    go mod tidy

    # Ensure the publish directory exists
    if (-Not (Test-Path $publishDirectory)) {
        New-Item -ItemType Directory -Path $publishDirectory | Out-Null
        Write-Host "Created publish directory: $publishDirectory"
    }

    # Attempt to build the module
    $buildSuccess = $false
    try {
        go build -o $publishDirectory -v ./...
        if ($LASTEXITCODE -eq 0) {
            $buildSuccess = $true

            Write-Host "Build succeeded for $modulePath"
        }
    } catch {
        Write-Warning "No main package to build or other errors occurred for $modulePath. Skipping..."
        
    }

    # Remove the directory if build failed or was skipped
    if (-Not $buildSuccess) {
        if (Test-Path $publishDirectory) {
            Remove-Item -Recurse -Force -Path "bin"
            Write-Warning "Removed publish directory: $publishDirectory"
        }
    }
    Pop-Location
}

Write-Host "Modules build files created successfully."
