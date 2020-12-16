include ./lib.mk
PROJECT_NAME=tcpproxy

export GOPATH:=$(HOME)/.gopath:$(PWD)
VERSION=`cat VERSION`

## build: compile the primary application
build: 
	@[ -d bin ] || mkdir bin
	go build -o bin/$(PROJECT_NAME) cmd/tcpproxy/main.go

## dist: compile the primary application for linux distribution
dist:
	@[ -d dist ] || mkdir dist
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dist/$(PROJECT_NAME) cmd/tcpproxy/main.go # && cp dist/$(PROJECT_NAME) docker/data-service/

## install-deps: fetch the go dependencies
install-deps:
	go get github.com/golang/lint/golint
	go get github.com/franela/goblin
	go get github.com/darrylwest/go-unique/unique
	go get gopkg.in/redis.v5
	go get github.com/darrylwest/cassava-logger/logger
	go get github.com/go-zoo/bone
	go get github.com/tidwall/gjson

## format: format the source files
format:
	( gofmt -s -w cmd/*.go )

## lint: lint the source files
lint:
	@( golint cmd/... )

## run : run the service
run:
	make build && ./bin/$(PROJECT_NAME)

## edit: edit all files
edit:
	vi -O cmd/tcpproxy/main.go proxy.go

