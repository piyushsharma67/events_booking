# Stage 1: Build the Go binary
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git and CA certificates for fetching dependencies
RUN apk add --no-cache git ca-certificates

# Copy go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN go build -o notification-service .

# Stage 2: Create a small image for running
FROM alpine:latest

WORKDIR /app

# Copy CA certificates for SMTP TLS
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /app/notification-service .

# Set environment variables defaults (optional)
ENV RABBITMQ_USER=guest \
    RABBITMQ_PASSWORD=guest \
    RABBITMQ_HOST=rabbitmq \
    RABBITMQ_PORT=5672 \
    RABBITMQ_QUEUE=notifications

# Run the service
CMD ["./notification-service"]
