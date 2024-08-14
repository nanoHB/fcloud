# Step 1: Use the official Golang image as the build environment
FROM golang:1.20-alpine AS builder

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Step 4: Copy the source code to the container
COPY . .

# Step 5: Build the Go application
RUN go build -o myapp

# Step 6: Use a minimal image to run the application
FROM alpine:latest

# Step 7: Set the working directory inside the container
WORKDIR /root/

# Step 8: Copy the built application from the builder stage
COPY --from=builder /app/myapp .

# Step 9: Copy the SQLite database file (if needed inside the container)
COPY --from=builder /app/myapp.db .

# Step 10: Expose the necessary port
EXPOSE 8080

# Step 11: Run the application
CMD ["./myapp"]
