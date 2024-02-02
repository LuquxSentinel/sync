build:
	@go build -o ./bin/sync

run: build
	@./bin/sync

test:
	@go test ./...