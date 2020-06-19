GO_VERSION = $(shell go version | cut -d' ' -f3)
SGH_VERSION = $(shell git describe --tags 2>/dev/null || git rev-parse --short HEAD)
DATE_FMT = +%Y-%m-%d
BUILD_DATE = $(shell date "$(DATE_FMT)")

LDFLAGS := -X main.SghVersion = $(SGH_VERSION) $(LDFLAGS)
LDFLAGS := -X main.GoVersion = $(GO_VERSION) $(LDFLAGS)
LDFLAGS := -X main.BuildDate = $(BUILD_DATE) $(LDFLAGS)

bin/sgh:
	@go build -trimpath -ldflags "$(LDFLAGS)" -o "$@" ./cmd
