build: 
	@go build -o ./bin/listener .

dev: build
	@./bin/listener

test:
	@go test -cover ./...