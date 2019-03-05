.PHONY: default test dependencies

default: test

dependencies:
	go get github.com/stretchr/testify/assert

test:
	go test -v ./...