# Use official Golang image as a build stage
FROM golang:1.23.5 AS builder

# Switch to root user explicitly
USER root

# Set the working directory inside the container
WORKDIR /app

# Copy the application source files
COPY . .

# If go.mod doesn't exist (fresh repo), initialize the module
RUN [ -f "go.mod" ] || go mod init cityweatherapi

# Download dependencies and tidy the module
RUN go mod tidy

# Expose application port
EXPOSE 8080

# Run the application
CMD ["go", "run", "server.go"]
