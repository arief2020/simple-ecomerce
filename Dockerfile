# Base image
FROM golang:1.22-alpine

# Set working directory
WORKDIR /app

# Install git
RUN apk add --no-cache git

# Copy files
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the application
RUN go build -o main ./app

# Expose port
EXPOSE 8000

# Run the application
CMD ["/app/main"]
