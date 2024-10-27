# Use the latest official Go image as the build environment
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency resolution
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application binary
RUN go build -o go-fiber-mvc

# Start a new stage with a minimal image
FROM gcr.io/distroless/base-debian11

# Set the working directory in the minimal container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/go-fiber-mvc .

# Copy the views folder for HTML templates
COPY --from=builder /app/views ./views

# Expose port 3000 for the Fiber app
EXPOSE 3000

# Run the application
CMD ["/app/go-fiber-mvc"]
