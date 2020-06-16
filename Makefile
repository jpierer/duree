run:
	go run *.go

watch:
	reflex --start-service -r '\.go$$' -- sh -c "go run *.go"

build:
	go get github.com/GeertJohan/go.rice
	go get github.com/GeertJohan/go.rice/rice
	rm -rf bin
	mkdir bin
	rice embed-go
	go build -o bin/duree