build:
	go build -o bin/hacktoberfest-2020

test:
	go test ./test/**

.PHONY: build test
