package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Item holds a simple url with its name
type Item struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Bookmarks holds multiple Bookmarks with a group name
type Bookmarks struct {
	Group string `json:"group"`
	Items []Item `json:"items"`
}

// readBookmarks reads given bookmarks json file
func (d *Duree) readBookmarks(filepath string) ([]Bookmarks, error) {
	if filepath == "" {
		filepath = "bookmarks.json"
	}

	jsonFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	jsonValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var bookmarks []Bookmarks
	json.Unmarshal(jsonValue, &bookmarks)

	return bookmarks, nil
}
