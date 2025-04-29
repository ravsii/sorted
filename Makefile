build:
	go build -o ./build/sorted ./cmd/sorted/main.go

test:
	go test -race ./...

.PHONY: build test
