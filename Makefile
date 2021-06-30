.PHONY: install test

test:
	go test ./...

install:
	go mod download
