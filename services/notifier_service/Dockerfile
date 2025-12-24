FROM golang:1.25-alpine

WORKDIR /app

# Install air
RUN go install github.com/air-verse/air@latest

# Copy go mod files first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Expose app port (change if needed)
EXPOSE 8004

# Run air
CMD ["air"]
