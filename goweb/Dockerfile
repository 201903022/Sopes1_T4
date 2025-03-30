# Etapa de compilación
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copiamos go.mod y descargamos dependencias (si las hubiera)
COPY go.mod ./
RUN go mod tidy

# Copiamos el resto del código fuente
COPY . .

# Compilamos el ejecutable
RUN go build -o goweb

# Imagen final mínima
FROM alpine:latest

# Instala certificados raíz para HTTPS (buena práctica)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copia el ejecutable desde la etapa anterior
COPY --from=builder /app/goweb .

# Exponer el puerto en el contenedor
EXPOSE 8080

# Comando de inicio del contenedor
CMD ["./goweb"]
