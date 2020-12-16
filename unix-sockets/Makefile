
all:
	make server client

client:
	go build -o bin/socket-client socket-client.go

server:
	go build -o bin/socket-server socket-server.go

format:
	gofmt -w *.go

.PHONY: format
.PHONY: test
.PHONY: watch
.PHONY: run
