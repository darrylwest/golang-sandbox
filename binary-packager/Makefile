
.SILENT:

all:
	echo "building pager..."
	go run generator.go && go build -o pager assets.go page.go

run:
	make all
	./pager

