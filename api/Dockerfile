FROM golang:1.21-alpine

WORKDIR /app

# Создаем go.mod
RUN go mod init api

# Копируем исходный код
COPY main.go .

# Добавляем зависимости
RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq

# Загружаем все зависимости
RUN go mod tidy

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Запуск приложения
EXPOSE 8000
CMD ["./main"] 