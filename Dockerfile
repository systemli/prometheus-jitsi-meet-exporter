FROM golang:1.17.0-alpine as builder

WORKDIR /go/src/github.com/systemli/prometheus-jitsi-meet-exporter

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

ADD . /go/src/github.com/systemli/prometheus-jitsi-meet-exporter
RUN apk --update add ca-certificates && \
    go get -d -v && \
    go mod download && \
    go mod verify && \
    CGO_ENABLED=0 go build -ldflags="-w -s" -o /prometheus-jitsi-meet-exporter


FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /prometheus-jitsi-meet-exporter /prometheus-jitsi-meet-exporter

USER appuser:appuser

EXPOSE 9888

ENTRYPOINT ["/prometheus-jitsi-meet-exporter"]
