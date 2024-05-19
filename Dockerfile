# Builder stage: This stage installs build tools and dependencies.
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /build

# Copy go mod and sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o exchangemonitor ./cmd/

# Final stage: This stage builds the final image with the compiled Go binary.
FROM alpine:3.14

# Set a working directory
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /build/exchangemonitor .

# Copy the configuration directory into the container
COPY config ./config

# Expose port 8080 to the outside world
EXPOSE 8080

# Environment variable to specify the configuration path
ENV CONFIG_PATH="/app/config"

# Command to run the executable
CMD ["./exchangemonitor"]
