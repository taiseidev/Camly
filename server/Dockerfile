# ベースイメージ
FROM golang:1.21-alpine

# Install required tools
RUN apk update && apk add --no-cache git curl

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY ./app/go.mod ./app/go.sum ./

# Download dependencies
RUN go mod download

# Copy the application source code
COPY app/ ./

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["go", "run", "main.go"]
