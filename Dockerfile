# Get Image from Docker Hub
FROM golang:1.22.5-bullseye

# Set working directory
WORKDIR /usr/app

# Copy go.mod and go.sum files first for better Docker caching
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the migrations folder into the container
COPY . .

# Install go-migrate CLI tool with PostgreSQL support
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Build the Go app
RUN go build -o bin/main ./cmd/api/main.go

# Expose the port the app will run on
EXPOSE 8000

# Command to run migrations and start the application
CMD ["/bin/sh", "-c", "/go/bin/migrate -path /usr/app/migrations -database 'postgres://user:password@localhost:5432/dbname?sslmode=disable' up && ./bin/main"]
