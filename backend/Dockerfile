# Stage 1: Build
FROM golang:1.25.4-alpine AS builder

WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fuente
COPY . .

# Compilar el binario
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/bin/main ./cmd/main.go

# Stage 2: Runtime
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copiar el binario compilado
COPY --from=builder /app/bin/main .
# Copiar el archivo de configuración
COPY --from=builder /app/internal/infrastructure/config/config.yaml ./internal/infrastructure/config/

EXPOSE 8080

CMD ["./main"]

