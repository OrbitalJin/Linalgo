run:
	@go run ./cmd/lib/main.go

build:
	@go build -o ./bin/linalgo ./cmd/lib/main.go 

test:
	@go test ./...