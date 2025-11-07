# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files (if they exist)
COPY go.* ./

# Copy source code
COPY *.go ./

# Build the application
RUN go build -o main .

# Run stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./main"]
