# Usa la imagen oficial de Go 1.23 como base
FROM golang:1.23-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de dependencias y el código
COPY go.mod go.sum ./
RUN go mod download

# Copia todo el código del proyecto
COPY . .

# Compila la aplicación asegurándote de que el binario se genere
RUN go build -o /app/main ./cmd/main.go

# Verifica que el binario se haya creado
RUN ls -la /app/

# Exponer el puerto
EXPOSE 8000

# Comando para ejecutar la aplicación
CMD ["/app/main"]