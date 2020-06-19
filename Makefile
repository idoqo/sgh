GO_VERSION = $(shell go version | cut -d' ' -f3)
SGH_VERSION = $(shell git describe --tags 2>/dev/null || git rev-parse --short HEAD)
DATE_FMT = +%Y-%m-%d
BUILD_DATE = $(shell date "$(DATE_FMT)")

LDFLAGS := -X command.SghVersion=$(SGH_VERSION) $(LDFLAGS)
LDFLAGS := -X command.GoVersion=$(GO_VERSION) $(LDFLAGS)
LDFLAGS := -X command.BuildDate=$(BUILD_DATE) $(LDFLAGS)

bin/sgh:
	@go build -trimpath -ldflags "$(LDFLAGS)" -o "$@" ./cmd

clean:
	rm -f ./bin/sgh
