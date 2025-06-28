run: build
	@./bin/fs

build:
	@go build -o bin/fs .

test:
	@go test ./p2p -v 