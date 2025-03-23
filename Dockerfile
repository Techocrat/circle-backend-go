# Use official Go image as base
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Install git (for go get) and build tools
RUN apk update && apk add --no-cache git

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . .

# Build the Go app
RUN go build -o main ./cmd/main.go

# Expose port
EXPOSE 8080

# Run the executable
CMD ["./main"]
