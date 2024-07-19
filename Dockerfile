# Etapa de construcción
FROM golang:1.22.2 AS build

# Configura el directorio de trabajo
WORKDIR /app

# Copia los archivos del proyecto
COPY . .

# Descarga las dependencias y compila el binario
RUN go mod tidy
RUN go build -o main ./cmd

# Etapa final
FROM alpine:latest

# Instala dependencias necesarias
RUN apk --no-cache add ca-certificates

# Crea un directorio para la aplicación
WORKDIR /root/

# Copia el binario desde la etapa de construcción
COPY --from=build /app/main .

# Define el comando de inicio
CMD ["./main"]


