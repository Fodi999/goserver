 # Используем официальный образ Go с поддержкой CGO
FROM golang:1.22.0

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем остальные файлы проекта
COPY . .

# Компилируем приложение
RUN CGO_ENABLED=1 GOOS=linux go build -o goserver ./cmd/goserver

# Указываем команду для запуска приложения
CMD ["./goserver"]
