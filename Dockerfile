# Stage 1: Build Stage
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy go.mod and go.sum first to leverage Docker's layer caching for dependencies
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application binary
RUN go build -o mini-social-media .

# Stage 2: Production Stage
FROM debian:bookworm-slim

# Set the working directory
WORKDIR /usr/src/app

# Copy the built binary from the builder stage
COPY --from=builder /usr/src/app/mini-social-media .

# Expose the application's port
EXPOSE 8000

# Command to run the executable
ENTRYPOINT ["./mini-social-media"]
