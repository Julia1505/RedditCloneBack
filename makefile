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

build:
	go build -o bin/redditclone cmd/redditclone/main.go
.PHONY:build

run: build
	./bin/redditclone

bd:
	docker compose up -d
.PHONY:bd


