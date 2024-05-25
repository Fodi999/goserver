param (
    [string]$projectDir = $(throw "Please specify a project directory")
)

# Проверка наличия git
if (-not (Get-Command git -ErrorAction SilentlyContinue)) {
    throw "Git is not installed or not in PATH"
}

# Проверка наличия go
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    throw "Go is not installed or not in PATH"
}

# Клонирование репозитория
Write-Host "Cloning repository..."
git clone https://github.com/Fodi999/goserver.git $projectDir
if ($LASTEXITCODE -ne 0) {
    throw "Failed to clone repository"
}

# Переход в каталог проекта
Set-Location -Path $projectDir

# Компиляция goinit.exe
Write-Host "Building goinit.exe..."
go build -o goinit.exe init.go
if ($LASTEXITCODE -ne 0) {
    throw "Failed to build goinit.exe"
}

# Запуск goinit.exe
Write-Host "Running goinit.exe..."
.\goinit.exe
if ($LASTEXITCODE -ne 0) {
    throw "Failed to run goinit.exe"
}

# Запуск сервера
Write-Host "Starting server..."
Set-Location -Path "cmd/goserver"
go run main.go
if ($LASTEXITCODE -ne 0) {
    throw "Failed to start server"
}
