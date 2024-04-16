build:
	@go build -o bin/passin ./cmd

run: build
	@./bin/passin

test:
	@go test -v ./...

.PHONY: build run test
