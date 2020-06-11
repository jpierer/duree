package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	listenAddr := os.Getenv("dureeListenAddr")
	if listenAddr == "" {
		listenAddr = "0.0.0.0:8080"
	}

	// bookmarkFile := os.Getenv("dureeBookmarksFile")
	bookmarkFile := "/Users/julianpierer/dev/duree/bookmarks.json"

	if bookmarkFile == "" {
		bookmarkFile = "~/.bookmarks.json"
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		os.Exit(1)
	}()

	duree := &Duree{listenAddr: listenAddr, bookmarkFilepath: bookmarkFile}
	duree.Run()
}
