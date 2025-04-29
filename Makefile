build:
	go build -o ./build/sorted ./cmd/sorted/*

install:
	go install ./cmd/sorted/sorted.go

test:
	go test -race ./...

.PHONY: build test
