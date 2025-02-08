build: 
	@go build -o ./bin/evm ./cmd/evm/main.go

dev: build
	@./bin/evm

test:
	@go test -cover -count=1 ./...

seed:
	@go run ./scripts/seed_address/main.go