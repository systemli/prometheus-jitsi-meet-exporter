# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
    
# Copy go mod and sum files
COPY . .
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Build the Go app
RUN go build -o prometheus-jitsi-meet-exporter .

USER appuser:appuser

# Expose port 8080 to the outside world
EXPOSE 9888

# Command to run the executable
ENTRYPOINT ["./prometheus-jitsi-meet-exporter"]