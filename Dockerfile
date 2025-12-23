# Tahap Build
FROM golang:1.25-alpine AS builder

# Set direktori kerja di dalam container
WORKDIR /app

# Salin file go.mod dan go.sum untuk men-download dependensi
COPY go.mod go.sum ./
RUN go mod download

# Salin seluruh source code aplikasi
COPY . .

# Build aplikasi Go
# -o ./main: Menentukan nama output binary adalah 'main'
# ./cmd/web/main.go: Path ke file main package
RUN CGO_ENABLED=0 GOOS=linux go build -o ./main ./cmd/web/main.go

# Tahap Final
FROM alpine:latest

# Install tzdata untuk informasi timezone
RUN apk add --no-cache tzdata

WORKDIR /app

# Salin binary yang sudah di-build dari tahap builder
COPY --from=builder /app/main .

# Salin direktori 'db' yang mungkin berisi file migrasi atau data seeder
COPY db ./db

# Expose port yang digunakan oleh aplikasi
EXPOSE 8080

# Perintah untuk menjalankan aplikasi
CMD ["./main"]
