# Create our container from a very lightweight image
FROM golang:alpine

# Set our working directory for the container
WORKDIR /app

# Copy our entire application into the Working Directory
COPY ./ /app

# Install our Go Modules
RUN go mod download

# Set the entry point for our application
ENTRYPOINT go run cmd/server.go