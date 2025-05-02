build:
	go build -o ./build/sorted ./cmd/sorted/*

install:
	go install ./cmd/sorted/sorted.go

test:
	go test -race ./...

lint:
	golangci-lint run ./...

help:
	go run ./cmd/sorted/sorted.go --help

.PHONY: build install test lint help
