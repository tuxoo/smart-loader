.PHONY: all lint build
.SILENT:

tidy:
	go mod tidy

clean:
	go clean -modcache

build-loader:
	go build -o ./.bin/app ./cmd/main.go

docker:
	docker compose up