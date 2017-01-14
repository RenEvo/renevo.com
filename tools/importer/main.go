package main

import (
	"bytes"
	"flag"
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

var minWordCount = flag.Int("min-words", 100, "The minimum amount of words to keep")

func main() {
	flag.Parse()

	output := path.Clean("./imported")
	parser := gofeed.NewParser()

	contentCount := 0

	feeds := []string{
		"http://renevo.com/blogs/developer/atom.aspx",
		"http://renevo.com/blogs/community_blogs/atom.aspx",
		"http://renevo.com/blogs/dotnet/atom.aspx",
	}
	for _, feed := range feeds {
		parsed, err := parser.ParseURL(feed)

		if err != nil {
			log.Fatalf("Failed to parse feed: %v", err)
		}

		log.Printf(parsed.Title)

		feedUrl, _ := url.Parse(feed)
		feedBasePath, _ := path.Split(feedUrl.Path)

		for _, item := range parsed.Items {
			if !strings.EqualFold(item.Author.Name, "tom anderson") {
				log.Printf("Skipping %v's post %s", item.Author.Name, item.Title)
				continue
			}

			postUrl, _ := url.Parse(item.Link)
			postPath, postFileName := path.Split(postUrl.Path)

			dir := path.Join(output, feedBasePath, fmt.Sprintf("%d", item.PublishedParsed.Year()), fmt.Sprintf("%d", item.PublishedParsed.Month()), fmt.Sprintf("%d", item.PublishedParsed.Day()))
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
			yamlFrontMatter += "original: " + item.Link + "\n"
			yamlFrontMatter += "file: " + postFileName + "\n"
			yamlFrontMatter += "path: " + postPath + "\n"
			yamlFrontMatter += "author: " + strings.ToLower(item.Author.Name) + "\n"

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

			bodyString := string(body)
			markerLocation := strings.LastIndex(string(bodyString), "![][")
			if markerLocation > 0 {
				wordCount := len(strings.Split(bodyString[:markerLocation], " "))
				yamlFrontMatter += "words: " + fmt.Sprintf("%v", wordCount) + "\n"

				if wordCount < *minWordCount {
					log.Printf("Post seemed unworthy of keeping: %v/%v words", wordCount, *minWordCount)
					continue
				}
			} else {
				log.Printf("Post didn't have any words worth mentioning %s", bodyString)
				continue
			}

			yamlFrontMatter += "---\n"

			if err := ioutil.WriteFile(filePrefix+".md", append([]byte(yamlFrontMatter), body...), 0644); err != nil {
				log.Printf("ERR: Failed to write file: %v", err)
				continue
			}

			contentCount++
		}
	}

	log.Printf("Finished Downloading %v items", contentCount)
}
