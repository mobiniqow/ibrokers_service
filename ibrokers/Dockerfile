# Step 1: Build the application
FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o /app/tour ./main.go

# Step 2: Run the application
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/tour /app/tour

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["./tour"]
