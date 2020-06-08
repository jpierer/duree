package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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

// read reads given bookmarks json file
func (d *Duree) read(filepath string) ([]Bookmarks, error) {
	if filepath == "" {
		return nil, errors.New("No Filepath given")
	}

	jsonValue, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var bookmarks []Bookmarks
	json.Unmarshal(jsonValue, &bookmarks)

	return bookmarks, nil
}

// write writes given bookmarks to json file
func (d *Duree) write(filepath string, bookmarks []Bookmarks) error {
	if filepath == "" {
		return errors.New("No Filepath given")
	}

	jsonValue, err := json.Marshal(bookmarks)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath, jsonValue, 0755)
	if err != nil {
		return err
	}

	return nil
}
