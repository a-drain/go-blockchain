build:
	go build -o bin/blockchain

test:
	go test -cover -v ./...

run: build
	./bin/blockchain