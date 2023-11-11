# Langkah 1: Tentukan gambar dasar yang ingin Anda gunakan
FROM golang:1.20-alpine as Build

# Langkah 4: Buat direktori kerja (misalnya /app) di dalam container
WORKDIR /app

# Langkah 5: Salin kode sumber aplikasi Golang Anda ke dalam container
COPY . .

# Langkah 6: Kompilasi aplikasi Golang
RUN go build -o main main.go


FROM alpine:latest

# Update package repository and install required packages
RUN apk update && \
    apk add --no-cache bash tzdata && \
    rm -rf /var/cache/apk/*

# Set the timezone
ENV TZ=Asia/Jakarta

# We need to copy the binary from the build image to the production image.
COPY --from=Build /app/main .

## This is the port that our application will be listening on.
EXPOSE 54321

# This is the command that will be executed when the container is started.
ENTRYPOINT ["./main"]