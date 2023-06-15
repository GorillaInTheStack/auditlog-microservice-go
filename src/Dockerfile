# Use the official Go image as the base image
FROM golang:1.20.5-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go mod and sum files to the working directory
COPY go.mod go.sum ./

# Download and cache Go modules then tidy them
RUN go clean -modcache
RUN go mod download
RUN go mod tidy

# Copy the source code from the current directory to the working directory inside the container
COPY . .

# Build microservice
RUN go build -o auditlog


# Start microservice
CMD ["./auditlog"]