.DEFAULT_GOAL := install

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

shadow: fmt
	shadow ./...
.PHONY:shadow

install: vet
	go install github.com/coci/odin
.PHONY:build

build:
	go build main.go
.PHONY:build
