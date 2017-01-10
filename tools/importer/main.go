package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/metal3d/go-slugify"
	"github.com/mmcdole/gofeed"
)

func main() {

	feed := "http://renevo.com/blogs/developer/atom.aspx"
	output := path.Clean("./imported")

	parser := gofeed.NewParser()

	parsed, err := parser.ParseURL(feed)

	if err != nil {
		log.Fatalf("Failed to parse feed: %v", err)
	}

	log.Printf(parsed.Title)

	for _, item := range parsed.Items {
		dir := path.Join(output, strings.ToLower(parsed.Title), fmt.Sprintf("%d", item.PublishedParsed.Year()), fmt.Sprintf("%d", item.PublishedParsed.Month()))
		os.MkdirAll(dir, 0644)

		slug := slugify.Marshal(item.Title, true)
		filePrefix := path.Join(dir, slug)

		// save html version
		log.Printf("Post: '%v' by '%v' posted on '%v'", item.Title, item.Author.Name, item.PublishedParsed)
		if err := ioutil.WriteFile(filePrefix+".html", []byte(item.Content), 0644); err != nil {
			log.Printf("ERR: Failed to write file: %v", err)
			continue
		}

		// create markdown version
		yamlFrontMatter := "---\n"
		yamlFrontMatter += "title: " + item.Title + "\n"
		yamlFrontMatter += "published: true\n"
		yamlFrontMatter += "date: " + fmt.Sprintf(item.PublishedParsed.String()) + "\n"
		yamlFrontMatter += "tags: imported " + strings.ToLower(strings.Join(item.Categories, " ")) + "\n"
		yamlFrontMatter += "---\n"

		req, err := http.NewRequest("POST", "http://heckyesmarkdown.com/go/", bytes.NewBuffer([]byte("html="+url.QueryEscape(item.Content))))
		if err != nil {
			log.Printf("Failed to create request: %v", err)
			continue
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Failed to make request: %v", err)
			continue
		}

		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err := ioutil.WriteFile(filePrefix+".md", append([]byte(yamlFrontMatter), body...), 0644); err != nil {
			log.Printf("ERR: Failed to write file: %v", err)
			continue
		}
	}
}
