PACKAGE=gocache
export GOPATH=$(HOME)/.gopath:$(PWD)

build: 
	@[ -d bin ] || mkdir bin
	( go build -o bin/$(PACKAGE) src/main.go )

build-linux:
	@[ -d docker ] || mkdir docker
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker/$(PACKAGE) src/main.go

install-deps:
	go get github.com/darrylwest/go-unique/unique
	go get github.com/golang/lint/golint
	go get github.com/franela/goblin
	go get -u github.com/darrylwest/cassava-logger/logger

format:
	( gofmt -s -w src/*.go src/*/*.go test/*.go )

lint:
	@( golint ./... && golint cache/... )

test:
	@( go vet src/*/*.go && go vet src/*.go && cd test && go test )
	@( make lint )

watch:
	clear && go-watcher --loglevel=4 --watch-folders=src,test

run:
	./docker/run.sh

edit:
	make format
	vi -O3 src/*/*.go test/*.go src/*.go

.PHONY: format lint test qtest watch run

