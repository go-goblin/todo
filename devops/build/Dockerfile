# Этап сборки
FROM golang:1.24-alpine AS builder

# Установим зависимости
RUN apk add --no-cache git

# Установим swag
RUN go install github.com/swaggo/swag/cmd/swag@latest
ENV PATH="/go/bin:$PATH"

# Рабочая директория
WORKDIR /app

# Копируем go.mod и go.sum отдельно для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Генерируем Swagger-документацию
RUN swag init -g cmd/server/main.go

# Сборка бинарника
RUN go build -o server cmd/server/main.go

# Финальный минимальный образ
FROM alpine:latest

# Установим сертификаты для HTTPS-запросов
RUN apk add --no-cache ca-certificates

# Копируем бинарник
COPY --from=builder /app/server /server

# Копируем Swagger-доки
COPY --from=builder /app/docs /docs

# Рабочая директория
WORKDIR /

# Порт (опционально)
EXPOSE 8080

# Запуск сервера
ENTRYPOINT ["/server"]
