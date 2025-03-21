FROM golang:1.21-alpine

WORKDIR /src/api

# Копируем файлы модуля
COPY go.mod go.sum* ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Проверяем модули и собираем приложение
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# Используем минимальный образ для запуска
FROM alpine:latest

WORKDIR /app

# Копируем собранное приложение
COPY --from=0 /app/main .

# Запускаем приложение
EXPOSE 8000
CMD ["./main"] 