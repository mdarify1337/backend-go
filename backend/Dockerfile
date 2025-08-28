# Use official Golang image
FROM golang:1.22

# Set working directory
WORKDIR /app

# Copy go mod files first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build Go app
RUN go build -o main .

# Expose backend port
EXPOSE 3000

# Run app
CMD ["./main"]
