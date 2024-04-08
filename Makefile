BUILDNUM             ?= $(GITLAB_BUILD_NUMBER)
GIT_SHA              ?= $(shell git describe --always --abbrev=8)
DEB_TARGET_ARCH      ?= armhf
ARCH                 ?= $(DEB_TARGET_ARCH)
TARGET               ?= osid
VERSION               = $(shell head -1 debian/changelog | sed 's/.*(\(.*\)).*/\1/')
PATH                 := /usr/local/go/bin:$(PATH)
PWD                   = $(shell pwd)

TARGETS               = build/$(ARCH)/$(TARGET)
PRODUCTION           ?= disable
PLATFORM             ?= platformMIL3

# GO
GOOS               ?= linux
HOST                = x86_64-linux-gnu
CC                  = $(HOST)-gcc
STRIP               = $(HOST)-strip
GO                  = GO111MODULE=on CGO_ENABLED=0 GOOS=$(GOOS) go
CGO                 = GO111MODULE=on CGO_ENABLED=1 GOOS=$(GOOS) go
ifeq ($(ARCH),armhf)
HOST                = arm-linux-gnueabihf
CC                  = $(HOST)-gcc
STRIP               = $(HOST)-strip
GO                 := GOARCH=arm GOARM=7 $(GO)
CGO                := CC=$(CC) GOARCH=arm GOARM=7 $(CGO)
endif

GOLDFLAGS          += -s -w
GOFLAGS             = -ldflags "$(GOLDFLAGS)" -mod=vendor

.PHONY: all osid test test_coverage

all: osid

all: $(TARGETS)

$(ARCH)/$(TARGET):
	mkdir -p build/$(ARCH)
	$(GO) build -buildmode=pie $(GOFLAGS) -tags "$(BUILD_ARCH_TAG) $(PLATFORM)" -o build/$@ ./cmd

build/$(ARCH)/$(TARGET): $(ARCH)/$(TARGET)
	if [ "$(PRODUCTION)" = "enable" ]; then \
		upx-ucl $@ ;\
	fi

test:
	(GO) test -v ./osi_test

test_coverage:
	go test -v -cover -coverpkg=./osi/...,./pkg/... -coverprofile=coverage.out ./osi_test
	go tool cover -func coverage.out
