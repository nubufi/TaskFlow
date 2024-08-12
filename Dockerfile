# Start from the official Golang image
FROM golang:1.23rc2-alpine3.19 AS builder

# Install necessary build tools
RUN apk add --no-cache git curl

# Install air for live reloading
RUN go install github.com/air-verse/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest
# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the application with air
CMD ["air"]

