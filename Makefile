.PHONY: default test dependencies generator metadata

default: test

dependencies:
	go get -t -v ./...

generator:
	go build -o bin/lcgenerator github.com/elliotcourant/gomonetary/lcgenerator

metadata: generator
	./bin/lcgenerator

test:
	go test -v ./...
	make metadata
	go test -v ./...

coverage:
	./coverage.sh