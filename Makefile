build: 
	@go build -o ./bin/app ./cmd/app

dev: build
	@./bin/app

test:
	@go test -cover -count=1 ./...