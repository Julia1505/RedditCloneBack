.DEFAULT_GOAL := run

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	glint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build hello.go
.PHONY:build

run: build
	./hello
