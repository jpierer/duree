run:
	go run *.go

watch:
	reflex --start-service -r '\.go$$' -- sh -c "go run *.go --listenAddr \"0.0.0.0:8080\" --bookmarkFile \"/Users/julianpierer/Library/Application Support/Google/Chrome/Default/Bookmarks\""

build:
	go get github.com/GeertJohan/go.rice
	go get github.com/GeertJohan/go.rice/rice
	rm -rf bin
	mkdir bin
	rice embed-go
	go build -o bin/duree

deploy:
	cp bin/duree /usr/local/bin
	