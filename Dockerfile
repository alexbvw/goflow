# Use a lightweight Go image.
FROM golang:1.21-alpine

# Install git, required for fetching Go dependencies.
RUN apk add --no-cache git

# Install Air for hot reloading.
RUN go install github.com/cosmtrek/air@latest

# Set the working directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum and download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your application's code.
COPY . .

# Expose the port your app runs on.
EXPOSE 3000

# Start the application with Air for hot reloading.
CMD ["air"]