GOCMD        = go
GOBUILD      = $(GOCMD) build
GOCLEAN      = $(GOCMD) clean
GOTEST       = $(GOCMD) test
GOVET        = $(GOCMD) vet
GOGET        = $(GOCMD) get
GOX          = $(GOPATH)/bin/gox
GOGET        = $(GOCMD) get

GIT_VERSION  := $(shell git --no-pager describe --tags --always)
GIT_COMMIT   := $(shell git rev-parse --verify HEAD)

GOX_ARGS     = -output="$(BUILD_DIR)/{{.Dir}}-${GIT_VERSION}-{{.OS}}-{{.Arch}}" -osarch="linux/amd64 linux/arm linux/arm64 darwin/amd64 freebsd/amd64"

BUILD_DIR    = build
BINARY_NAME  = prometheus-jitsi-meet-exporter

all: clean vet test build

build:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v

vet:
	${GOVET} ./...

test:
	${GOTEST} ./...

coverage:
	${GOTEST} -coverprofile=coverage.txt -covermode=atomic ./...

clean:
	$(GOCLEAN)
	rm -f $(BUILD_DIR)/*

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

release:
	${GOGET} -u github.com/mitchellh/gox
	${GOX} -ldflags "${LD_FLAGS}" ${GOX_ARGS}

docker:
	docker build --rm --force-rm --no-cache -t systemli/prometheus-jitsi-meet-exporter .

.PHONY: all vet test coverage clean build run release docker
