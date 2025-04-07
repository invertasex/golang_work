# Используем официальный образ Go
FROM golang:1.20

WORKDIR /app

# Копируем go.mod и go.sum для кеширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной код
COPY . .

# Сборка
RUN go build -o grpc-service main.go

# Запуск
EXPOSE 50051
CMD ["./grpc-service"]
