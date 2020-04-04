# Jitsi Meet Metrics Exporter

![Integration](https://github.com/systemli/prometheus-jitsi-meet-exporter/workflows/Integration/badge.svg?branch=master) ![Quality](https://github.com/systemli/prometheus-jitsi-meet-exporter/workflows/Quality/badge.svg?branch=master) [![Docker Automated build](https://img.shields.io/docker/automated/systemli/prometheus-jitsi-meet-exporter.svg)](https://hub.docker.com/r/systemli/prometheus-jitsi-meet-exporter/) [![MicroBadger Size](https://img.shields.io/microbadger/image-size/systemli/prometheus-jitsi-meet-exporter.svg)](https://hub.docker.com/r/systemli/prometheus-jitsi-meet-exporter/)

Prometheus Exporter for Jitsi Meet written in Go. Based on [Jitsi Meet Exporter](https://git.autistici.org/ai3/tools/jitsi-prometheus-exporter) from [Autistici](https://www.autistici.org/)

## Usage

```
go get github.com/systemli/prometheus-jitsi-meet-exporter
go install github.com/systemli/prometheus-jitsi-meet-exporter
$GOPATH/bin/prometheus-jitsi-meet-exporter
```

### Docker

```
docker run -p 9888:9888 systemli/prometheus-jitsi-meet-exporter:latest -videobridge-url http://jitsi:8888/stats 
```

## License

GPLv3
