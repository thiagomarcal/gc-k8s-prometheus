# Use the official Golang image to build the application
FROM golang:1.23.2 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the application
# RUN go build -o server .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

# Use a minimal base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

RUN chmod +x server

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./server"]
