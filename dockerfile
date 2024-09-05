# Step 1: Use the official Golang image as the build environment
FROM golang:1.23-alpine AS builder

# Step 2: Set the working directory inside the container
WORKDIR /app
RUN go install github.com/air-verse/air@latest

# Step 3: Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Step 4: Copy the source code to the container
COPY . .

# Step 10: Expose the necessary port
EXPOSE 8080

# Step 11: Run the application
CMD ["air", "-c", ".air.toml"]
