# Build Stage
FROM golang:alpine AS build

WORKDIR /app

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o prometheus-jitsi-meet-exporter .

# Final Stage
FROM alpine

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

COPY --from=build /app/prometheus-jitsi-meet-exporter .

USER appuser:appuser

EXPOSE 9888

ENTRYPOINT ["./prometheus-jitsi-meet-exporter"]