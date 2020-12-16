export GOPATH:=$(HOME)/.gopath:$(PWD)
TARGET=/usr/local/bin

build: 
	@[ -d bin ] || mkdir bin
	( /bin/rm -f bin/* )
	( go build -o bin/pubsub src/main.go )

build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o linux/unique src/main.go

docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o linux/unique-tcp src/unique-tcp.go
	( cd linux && ./build.sh )

format:
	( gofmt -s -w src/*.go src/unique/*.go test/*.go )

lint:
	@( golint src/... && golint test/... )

test:
	@( go vet src/unique/*.go && go vet src/unique/*.go && go vet src/*.go && cd test/ && go test -cover )
	@( make lint )

run:
	go run src/main.go

watch:
	go-watcher

edit:
	vi -O3 src/*/*.go test/*.go src/*.go

.PHONY: format test watch examples
