# ---------- build ----------
FROM golang:1-alpine AS builder

WORKDIR /app

# кеш зависимостей
COPY go.mod go.sum ./
RUN go mod download

# копируем код
COPY . .

# сборка
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o server

# ---------- runtime ----------
FROM alpine:3.20

WORKDIR /app

# добавляем certs
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./server"]