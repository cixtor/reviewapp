run: build
	./server

build:
	go build -o server -- src/*.go
