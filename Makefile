build:
	go build -o bin/blockchain

test:
	go test ./...

run: build
	./bin/blockchain