package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	listenAddr := flag.String("listenAddr", "", "host:port")
	bookmarkFile := flag.String("bookmarkFile", "", "path to the bookmarks file of your browser")

	flag.Parse()

	// check for required parameters
	if *listenAddr == "" || *bookmarkFile == "" {
		fmt.Print("\nUsage of ./duree:\n\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		os.Exit(1)
	}()

	duree := &Duree{listenAddr: *listenAddr, bookmarkFilepath: *bookmarkFile}
	duree.Run()
}
