
all:
	go build -o bin/chat ./chat

install-deps:
	go get github.com/dgrijalva/jwt-go
	go get -u github.com/cosmos72/gomacro
	go get -u github.com/integrii/flaggy
	go get github.com/disintegration/imaging
	go get github.com/fogleman/gg
	go get -u github.com/disintegration/gift
	go get gonum.org/v1/plot/...
	go get -u github.com/gobuffalo/packr/...
	go get -u github.com/tidwall/buntdb
	go get github.com/golang/freetype
	go get -u github.com/shurcooL/vfsgen
	go get github.com/francoispqt/gojay
	go get github.com/thedevsaddam/govalidator
	go get -u github.com/shuLhan/go-bindata/...
	go get github.com/eclipse/paho.mqtt.golang
	go get github.com/briandowns/simple-httpd
	go get github.com/yanatan16/itertools
	go get github.com/hashicorp/consul/api
	go get github.com/buger/goterm
	go get -u github.com/darrylwest/spot-cache/spotcache
	go get -u github.com/darrylwest/go-unique/unique
	go get github.com/codegangsta/negroni
	go get github.com/pborman/uuid
	go get -u github.com/darrylwest/cassava-logger/logger
	go get github.com/franela/goblin
	go get gopkg.in/redis.v5
	go get golang.org/x/crypto/nacl/secretbox
	go get golang.org/x/crypto/nacl/box
	go get golang.org/x/crypto/scrypt
	go get github.com/agl/ed25519
	go get github.com/Pallinder/go-randomdata
	go get -u github.com/coreos/bbolt/...
	go get github.com/br0xen/boltbrowser
	go get github.com/docker/docker/client
	go get github.com/docker/docker/api
	go get github.com/nats-io/go-nats
	go get github.com/spf13/cobra
	go get -u github.com/gizak/termui
	go get github.com/go-zoo/bone
	go get github.com/fsnotify/fsnotify
	go get github.com/jung-kurt/gofpdf
	go get github.com/google/tcpproxy
	go get -u github.com/tidwall/gjson
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go

format:
	( cd socket-client-server ; gofmt -s -w *.go )
	( cd web-sockets ; gofmt -s -w *.go )
	( cd chat ; gofmt -s -w *.go )
	( cd util ; gofmt -s -w *.go )

test:
	( cd trace; go test -cover )

watch:
	./watcher.js

run:
	./bin/chat -addr=":3000"

.PHONY: format
.PHONY: test
.PHONY: watch
.PHONY: run
