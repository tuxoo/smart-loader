.PHONY: all lint test build
.SILENT:

lint:
	golangci-lint run

tidy:
	go mod tidy

clean:
	go clean -modcache

build:
	go build -o ./.bin/app ./cmd/idler-service/main.go

docker:
	docker compose up