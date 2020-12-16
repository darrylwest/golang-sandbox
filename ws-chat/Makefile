
run:
	go run *.go

build:
	go-bindata -o data.go home.html favicon.ico
	go build -o chat-service *.go

docker-build:
	go-bindata -o data.go home.html favicon.ico
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o chat-service *.go
