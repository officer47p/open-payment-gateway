build: 
	@go build -o ./bin/app .

dev: build
	@./bin/app

test:
	@go test -cover -count=1 ./...