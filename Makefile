.PHONY: run

install:
    go mod tidy

build:
    go build -o build/server cmd/main.go

run:
    go run cmd/main.go