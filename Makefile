.DEFAULT_GOAL := build

APP_NAME=monitor-client
APP_BINARY=bin/$(APP_NAME)
APP_BINARY_UNIX=bin/$(APP_NAME)_unix_amd64

all: build

.PHONY: test
test: ## test
	go test -v ./...

.PHONY: build
build: ## build
	CGO_ENABLED=0 go build -o $(APP_BINARY) -v *.go


.PHONY: build-linux
build-linux: ## build linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(APP_BINARY_UNIX) -v *.go


