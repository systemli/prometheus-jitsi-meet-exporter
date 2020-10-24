FROM golang:1.15.3-alpine
ENV GO111MODULE=on
WORKDIR /go/src/github.com/systemli/prometheus-jitsi-meet-exporter
ADD . /go/src/github.com/systemli/prometheus-jitsi-meet-exporter
RUN go build -o /prometheus-jitsi-meet-exporter

FROM alpine
WORKDIR /app
COPY --from=0 /prometheus-jitsi-meet-exporter /prometheus-jitsi-meet-exporter

EXPOSE 9888
ENTRYPOINT ["/prometheus-jitsi-meet-exporter"]
