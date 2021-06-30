.PHONY: install lint test

default: lint test

install:
	@go version &&\
    go get -u golang.org/x/lint/golint

lint:
	@cd examples && \
	go vet ./... &&\
	golint ./...


test:
	@cd examples && \
	go test ./...
