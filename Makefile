.PHONY: install lint test

default: lint, test

install:
	@go version &&\
    go get -u golang.org/x/lint/golint

lint:
	@go vet ./... &&\
	golint ./...


test:
	@go test ./...
