FROM alpine:3.18.0 as builder

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


FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY prometheus-jitsi-meet-exporter /prometheus-jitsi-meet-exporter

USER appuser:appuser

EXPOSE 9888

ENTRYPOINT ["/prometheus-jitsi-meet-exporter"]
