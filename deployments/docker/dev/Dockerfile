# Official Golang image
FROM golang:1.20

# Set workdir
WORKDIR /app

# Copy go.mod and go.sum files from the root parent directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY ../.. .

# Build the app (not required for tests, but included for consistency)
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -v -o cook cmd/cook/main.go

# Port
EXPOSE 8080

# Set the entrypoint to run the test
ENTRYPOINT ["./cook"]

