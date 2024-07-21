# Start from the official golang image with Alpine Linux as a builder stage
FROM golang:1.21.10-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY ../go.mod ../go.sum ./

# Download dependencies using go mod
RUN go mod download

# Copy the entire source code to the working directory
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Start a new stage from the Alpine Linux image
FROM alpine:latest

# Set the working directory in the new stage
WORKDIR /root/

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
