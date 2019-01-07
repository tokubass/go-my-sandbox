GOOS      = $(shell go env GOOS)
GOARCH    = $(shell go env GOARCH)
GOHOSTOS  = $(shell go env GOHOSTOS)
GOHOSTARCH = $(shell go env GOHOSTARCH)
VERSION   = $(shell git rev-parse HEAD)
BUILD_AT  = $(shell date '+%Y/%m/%d %H:%M:%S')
LDFLAGS   = -ldflags "-X \"main.Version=$(VERSION)\" -X \"main.BuildTime=$(BUILD_AT)\"  "
GOBUILD   = go build -v $(LDFLAGS)
CURRENT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

.PHONY: test
test:
	go test --tags 'test' -v ./...
