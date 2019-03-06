.PHONY: default tests dependencies generator metadata

default: tests

dependencies:
	go get -t -v ./...

generator:
	go build -o bin/lcgenerator github.com/elliotcourant/gomonetary/lcgenerator

metadata: generator
	./bin/lcgenerator

tests:
	go test -v ./...

coverage:
	./coverage.sh