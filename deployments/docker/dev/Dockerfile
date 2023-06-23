# Start from Alpine image with Go 1.20.3
FROM golang:1.20.5-alpine3.17 AS builder

# Set workdir
WORKDIR /go/src/github.com/foorester/cook

# Copy go.mod and go.sum from the top parent directory
COPY . .

# Download dependencies
RUN go mod download

# Build the app
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -v -o /go/bin/cook cmd/cook/main.go

# Use a minimal image as the final base
FROM alpine:3.17
#WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /go/bin/cook ./cook

# Port
EXPOSE 8080

# Entrypoint
ENTRYPOINT ["./cook"]
