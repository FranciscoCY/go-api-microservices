# -------- BUILD STAGE --------
FROM golang:1.25.1 AS builder

WORKDIR /app

# Copiar go.mod y go.sum primero (mejora caching de dependencias)
COPY go.mod go.sum ./
RUN go mod download

# Copiar todo el código
COPY cmd/ cmd/
COPY internal/ internal/
COPY proto/ proto/
COPY config/ config/

# Compilar el binario
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/server

# -------- RUNTIME STAGE --------
FROM alpine:latest

WORKDIR /app

# Copiar solo el binario desde la etapa builder
COPY --from=builder /app/app .

# Copiar configuración (si quieres tener config.yaml dentro del contenedor)
COPY config/config.yaml ./config/config.yaml

# Puerto gRPC
EXPOSE 50051

CMD ["./app"]
