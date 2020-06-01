run:
	go run *.go

build:
	mkdir bin; rice embed-go; go build -o bin/duree

watch:
	reflex --start-service -r '\.go$$' -- sh -c "go run *.go"
