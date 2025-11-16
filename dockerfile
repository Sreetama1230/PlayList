# Stage 1: Build Go binary
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o playlist-app

# Stage 2: Small runtime image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/playlist-app .
EXPOSE 8080
CMD ["./playlist-app"]
