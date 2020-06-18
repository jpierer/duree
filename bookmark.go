package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Item holds a simple url with its name
type Item struct {
	Title string
	URL   string
}

// Bookmarks holds multiple Bookmarks with a group name
type Bookmarks struct {
	Group string
	Items []Item
}

type ChromeBookmarks struct {
	Checksum string `json:"checksum"`
	Roots    struct {
		BookmarkBar struct {
			Children []struct {
				Children []struct {
					Children []struct {
						DateAdded string `json:"date_added"`
						GUID      string `json:"guid"`
						ID        string `json:"id"`
						Name      string `json:"name"`
						Type      string `json:"type"`
						URL       string `json:"url"`
					} `json:"children"`
					DateAdded    string `json:"date_added"`
					DateModified string `json:"date_modified"`
					GUID         string `json:"guid"`
					ID           string `json:"id"`
					Name         string `json:"name"`
					Type         string `json:"type"`
				} `json:"children"`
				DateAdded    string `json:"date_added"`
				DateModified string `json:"date_modified"`
				GUID         string `json:"guid"`
				ID           string `json:"id"`
				Name         string `json:"name"`
				Type         string `json:"type"`
			} `json:"children"`
		} `json:"bookmark_bar"`
	} `json:"roots"`
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

	var chromeBookmarks ChromeBookmarks
	bookmarks := []Bookmarks{}
	json.Unmarshal(jsonValue, &chromeBookmarks)

	// run through bookmark folders
	folders := chromeBookmarks.Roots.BookmarkBar.Children
	for i := 0; i < len(folders); i++ {
		if folders[i].Type != "folder" || folders[i].Name != "duree" {
			continue
		}

		dureeFolder := folders[i]

		// run through duree group folders
		for g := 0; g < len(dureeFolder.Children); g++ {

			group := dureeFolder.Children[g]
			items := []Item{}

			// get bookmark links
			for l := 0; l < len(group.Children); l++ {
				items = append(items, Item{Title: group.Children[l].Name, URL: group.Children[l].URL})
			}

			bookmarks = append(bookmarks, Bookmarks{Group: group.Name, Items: items})
		}
	}

	return bookmarks, nil
}
