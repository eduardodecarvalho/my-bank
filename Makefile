build:
	@go build -o bin/my-bank

run: build
	@./bin/my-bank

test:
	@go test -v ./...
