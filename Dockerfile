# Использование официального образа Go с поддержкой CGO
FROM golang:1.20

# Установка gcc
RUN apt-get update && apt-get install -y gcc

# Создание рабочей директории
WORKDIR /app

# Копирование всех файлов проекта в контейнер
COPY . .

# Установка зависимостей
RUN go mod download

# Компиляция goinit
RUN CGO_ENABLED=1 go build -o goinit init.go

# Команда по умолчанию для запуска сервера
CMD ["go", "run", "cmd/goserver/main.go"]


