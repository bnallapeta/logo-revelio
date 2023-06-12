# Use an official Golang runtime as the base image
FROM quay.io/projectquay/golang:1.20

# Install necessary dependencies for CGO
RUN dnf install -y gcc glibc-devel

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files into the container
COPY go.mod go.sum ./

# Download and cache the Go modules
RUN go mod download

# Copy the rest of the application source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=1 go build -o logo-revelio cmd/logo-revelio/main.go

# Expose the port that the web application listens on
EXPOSE 8080

# Define the command to run the application
CMD ["./logo-revelio"]
