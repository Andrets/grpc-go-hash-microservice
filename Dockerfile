# Используем официальный образ Golang как базовый
FROM golang:1.22.1-alpine3.19 as builder

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Копируем go.mod и go.sum в рабочую директорию
COPY go.mod go.sum ./

# Загружаем все зависимости
RUN go mod download

# Копируем исходный код в рабочую директорию
COPY . .

# Собираем приложение
RUN go build -o grpc-go-hash-microservice ./cmd/main.go

# Начинаем новую стадию сборки на основе Alpine
FROM alpine:3.19 as production

# Устанавливаем ca-certificates
RUN apk add --no-cache ca-certificates

# Копируем исполняемый файл из стадии builder в корень файловой системы контейнера
COPY --from=builder /app/grpc-go-hash-microservice /grpc-go-hash-microservice

# Запускаем приложение
CMD ["/grpc-go-hash-microservice"]

# Открываем порт 50051
EXPOSE 50051