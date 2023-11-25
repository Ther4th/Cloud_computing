FROM cgr.dev/chainguard/go:latest as builder

# Set the working directory inside the container
WORKDIR /app

COPY . .

# Download dependencies
RUN go mod download

# Build the Go application
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o main .

RUN mkdir /app/data
# Create a minimal final image
FROM scratch

# Copy the compiled application binary from the builder image
COPY --from=builder /app/main /app/main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]
