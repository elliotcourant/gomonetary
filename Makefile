.PHONY: default test

default: test

test:
	go test -v ./...