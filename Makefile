BINARY := $(shell basename "$(PWD)")
SOURCES := ./
GIT_COMMIT := $(shell git rev-list -1 HEAD)
GIT_VERSION := $(shell git describe --tags --abbrev=0)

PKG := $(shell cat go.mod | sed -n "s/^module \(.*\)$$/\1/p" | sed "s/terraform-provider-//g")
VER := $(shell git describe --tags --abbrev=0 | sed -n "s/^v\(.*\)$$/\1/p")
OS := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)

PLUGINS_DIR := ~/.terraform.d/plugins/${PKG}/$(VER)/$(OS)_$(ARCH)/

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.DEFAULT_GOAL := help

## build: Build the command line tool
build: clean
	CGO_ENABLED=0 go build \
	-ldflags '-w -extldflags "-static" -X main.gitCommit=$(GIT_COMMIT)' \
	-o ${BINARY} ${SOURCES}

## pack: Shrink the binary size
pack: build 
	upx -9 ${BINARY}

## install: Install this plugin
install: build
	mkdir -p ${PLUGINS_DIR}
	mv ${BINARY} ${PLUGINS_DIR}

## test: Starts unit test
test:
	go test -v ./... -coverprofile coverage.out

## clean: Clean the binary
clean:
	rm -f $(BINARY)