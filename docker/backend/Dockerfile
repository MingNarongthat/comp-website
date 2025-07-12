# Build stage
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Debug: Print contents of /app after build
RUN ls -la

# Final stage
FROM alpine:latest

WORKDIR /app

# Debug: Print contents of /app before copying
RUN ls -la

# Copy the binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Debug: Print contents of /app after copying
RUN ls -la

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./main"]
