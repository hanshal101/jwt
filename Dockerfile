# Use an official Go runtime as a base image
FROM golang:1.21.5

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any required dependencies
RUN go mod download

# Build the Go application
RUN go build -o jwt

# Expose port 9876 to the outside world
EXPOSE 9876

# Command to run the executable
CMD ["./jwt"]
