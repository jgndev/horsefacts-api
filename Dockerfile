############################################################
# First stage
############################################################
# Start from the latest Golang base image
FROM golang:1.20-buster as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
# Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

############################################################
# Second stage #
############################################################
FROM debian:buster-slim

WORKDIR /app

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /app/main .

# Install certificates
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by ‘go install’
ENTRYPOINT ["/app/main"]