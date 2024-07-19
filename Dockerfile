# Usa una imagen base de Go para la etapa de construcción
FROM golang:1.22.2-alpine AS builder

# Establece el directorio de trabajo
WORKDIR /app

# Copia el go.mod y go.sum y descarga las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código fuente
COPY . .

# Compila la aplicación
RUN go build -o main ./cmd/main.go

# Usa una imagen base más pequeña para el runtime
FROM alpine:latest

# Instala dependencias necesarias
RUN apk --no-cache add ca-certificates

# Crea un directorio para la aplicación
WORKDIR /root/

# Copia el binario compilado desde la etapa de construcción
COPY --from=builder /app/main .

# Establece el comando por defecto
CMD ["./main"]
