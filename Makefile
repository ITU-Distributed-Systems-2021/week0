.PHONY: install test

test: install
	go test ./...

install:
	go mod download
