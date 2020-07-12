package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1594527248, 0),

		Content: string("<!DOCTYPE html>\n<html lang=\"en\" data-theme=\"light\">\n\n    <head>\n        <meta charset=\"UTF-8\">\n        <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n        <title>duree</title>\n        <link rel=\"stylesheet\" href=\"/static/app.css\">\n    </head>\n\n    <body>\n\n        <div class=\"links-wrapper\">\n            {{range .Bookmarks}}\n            <div class=\"{{.Group}}-links group-container\">\n                <h1 data-group=\"{{.Group}}\">{{.Group}}</h1>\n                {{range .Items}}\n                <a class=\"link\"\n                    href=\"{{.URL}}\">{{.Title}}</a>\n                {{end}}\n            </div>\n            {{end}}\n\n        </div>\n\n    </body>\n\n</html>"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1589918547, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "index.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`templates`, &embedded.EmbeddedBox{
		Name: `templates`,
		Time: time.Unix(1589918547, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"index.html": file2,
		},
	})
}

func init() {

	// define files
	file4 := &embedded.EmbeddedFile{
		Filename:    "app.css",
		FileModTime: time.Unix(1594527644, 0),

		Content: string(":root {\n    --font-color: #202125;\n    --background-color: #fff;\n}\n\n@media (prefers-color-scheme: dark) {\n    :root {\n        --font-color: #e2e2e2;\n        --background-color: #202125;\n    }\n}\n\n* {\n    margin: 0;\n    padding: 0;\n    box-sizing: border-box;\n    font-family: \"Inter\";\n    color: var(--background-color);\n    font-size: 13.5px;\n}\n\nbody {\n    height: 100vh;\n    width: 100vw;\n    background-color: var(--background-color);\n    display: flex;\n    flex-direction: column;\n    justify-content: center;\n    align-items: center;\n}\n\nh1 {\n    font-weight: 600;\n    margin-bottom: 10px;\n    color: var(--font-color);\n    text-transform: lowercase;\n}\n\n.links-wrapper {\n    display: flex;\n}\n\n.links-wrapper div {\n    display: flex;\n    flex-direction: column;\n    margin-right: 110px;\n}\n\n.links-wrapper div:last-child {\n    margin-right: 0;\n}\n\n.links-wrapper div a {\n    text-decoration: none;\n    color: var(--font-color);\n    opacity: 0.5;\n    margin-bottom: 8px;\n    text-transform: lowercase;\n}"),
	}

	// define dirs
	dir3 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1592496230, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file4, // "app.css"

		},
	}

	// link ChildDirs
	dir3.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`static`, &embedded.EmbeddedBox{
		Name: `static`,
		Time: time.Unix(1592496230, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir3,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"app.css": file4,
		},
	})
}
