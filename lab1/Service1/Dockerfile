FROM cgr.dev/chainguard/go:latest as builder

# Set the working directory inside the container
WORKDIR /app

COPY . .

# Download dependencies
RUN go mod download

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Create a minimal final image
FROM scratch

# Copy the compiled application binary from the builder image
COPY --from=builder /app/main /app/main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]
