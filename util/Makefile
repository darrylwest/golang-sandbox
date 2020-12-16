vars:
	go run -ldflags "-X main.xyz=this-is-my-ldflags-test" vars.go

format:
	gofmt -s -w *.go

test:
	go vet *.go

watch:
	./watcher.js

run:
	./bin/chat -addr=":3000"

.PHONY: format
.PHONY: test
.PHONY: watch
.PHONY: run
