.PHONY: default test dependencies generator metadata

default: test

dependencies:
	go get github.com/stretchr/testify/assert
	go get github.com/kataras/golog

generator:
	go build -o bin/lcgenerator github.com/elliotcourant/gomonetary/lcgenerator

metadata: generator
	./bin/lcgenerator

test:
	go test -v ./...
	make metadata
	go test -v ./...