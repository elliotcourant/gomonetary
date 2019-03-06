.PHONY: default test dependencies

default: test

dependencies:
	go get github.com/stretchr/testify/assert

generator:
	go build -o bin/lcgenerator github.com/elliotcourant/gomonetary/lcgenerator

metadata: generator
	./bin/lcgenerator

test:
	go test -v ./...