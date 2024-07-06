.PHONY: run

install:
    go mod tidy

build:
    go build -o build/server main.go

run:
    go run main.go

docker:
	docker-compose up -d --build