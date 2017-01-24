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

	"sync/atomic"

	"time"

	"github.com/metal3d/go-slugify"
	"github.com/mmcdole/gofeed"
)

var minWordCount = flag.Int("min-words", 100, "The minimum amount of words to keep")

func main() {
	flag.Parse()

	blogFeeds := []string{
		"http://renevo.com/blogs/developer/atom.aspx",
		"http://renevo.com/blogs/community_blogs/atom.aspx",
		"http://renevo.com/blogs/dotnet/atom.aspx",
	}

	blogImporter := NewBlogImporter(&BlogImporterConfig{"./imported", blogFeeds, *minWordCount})

	blogImporter.Parse()

	log.Printf("Finished Downloading %v items", blogImporter.PostCount())
}

type BlogImporterConfig struct {
	OutputDir    string
	Feeds        []string
	MinWordCount int
}

type BlogImporter struct {
	parser    *gofeed.Parser
	postCount uint32
	config    *BlogImporterConfig
}

func NewBlogImporter(config *BlogImporterConfig) *BlogImporter {
	config.OutputDir = path.Clean(config.OutputDir)

	return &BlogImporter{
		parser: gofeed.NewParser(),
		config: config,
	}
}

func (this *BlogImporter) Parse() {
	completion := make(chan string)
	defer close(completion)
	now := time.Now()

	for _, feed := range this.config.Feeds {
		go func(thisFeed string) {
			this.parseFeed(thisFeed)
			completion <- thisFeed
		}(feed)
	}

	for i := 0; i < len(this.config.Feeds); i++ {
		finished := <-completion
		log.Printf("Finished: %s", finished)
	}

	log.Printf("Parse took: %v", time.Since(now))
}

func (this *BlogImporter) PostCount() uint32 {
	return this.postCount
}

func (this *BlogImporter) parseFeed(feed string) {
	now := time.Now()

	parsed, err := this.parser.ParseURL(feed)

	if err != nil {
		log.Fatalf("Failed to parse feed: %v", err)
	}

	log.Printf(parsed.Title)

	feedUrl, _ := url.Parse(feed)
	feedBasePath, _ := path.Split(feedUrl.Path)
	completion := make(chan *gofeed.Item)
	defer close(completion)

	for _, item := range parsed.Items {
		go func(thisItem *gofeed.Item) {
			now := time.Now()
			if err := this.parseFeedItem(feedBasePath, thisItem); err != nil {
				log.Print(err.Error())
			} else {
				atomic.AddUint32(&this.postCount, 1)
			}
			log.Printf("Parse Feed Item %s took: %v", thisItem.Title, time.Since(now))
			completion <- thisItem
		}(item)
	}

	for i := 0; i < len(parsed.Items); i++ {
		<-completion
	}

	log.Printf("Parse Feed %s took: %v", feed, time.Since(now))
}
func (this *BlogImporter) parseFeedItem(feedBasePath string, item *gofeed.Item) error {
	if !strings.EqualFold(item.Author.Name, "tom anderson") {
		return fmt.Errorf("Skipping %v's post %s", item.Author.Name, item.Title)
	}

	postUrl, _ := url.Parse(item.Link)
	postPath, postFileName := path.Split(postUrl.Path)

	dir := path.Join(this.config.OutputDir, feedBasePath, fmt.Sprintf("%d", item.PublishedParsed.Year()), fmt.Sprintf("%d", item.PublishedParsed.Month()), fmt.Sprintf("%d", item.PublishedParsed.Day()))
	os.MkdirAll(dir, 0644)

	slug := slugify.Marshal(item.Title, true)
	filePrefix := path.Join(dir, slug)

	// save html version
	log.Printf("Post: '%v' by '%v' posted on '%v'", item.Title, item.Author.Name, item.PublishedParsed)
	if err := ioutil.WriteFile(filePrefix+".html", []byte(item.Content), 0644); err != nil {
		return fmt.Errorf("ERR: Failed to write file: %v", err)
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

	bodyString, err := renderToMarkdown(item.Content)
	if err != nil {
		return err
	}

	markerLocation := strings.LastIndex(string(bodyString), "![][")
	if markerLocation > 0 {
		wordCount := len(strings.Split(bodyString[:markerLocation], " "))
		yamlFrontMatter += "words: " + fmt.Sprintf("%v", wordCount) + "\n"

		if wordCount < this.config.MinWordCount {
			return fmt.Errorf("Post seemed unworthy of keeping: %v/%v words", wordCount, this.config.MinWordCount)
		}
	} else {
		return fmt.Errorf("Post didn't have any words worth mentioning %s", bodyString)
	}

	yamlFrontMatter += "---\n"

	if err := ioutil.WriteFile(filePrefix+".md", append([]byte(yamlFrontMatter), []byte(bodyString)...), 0644); err != nil {
		return fmt.Errorf("ERR: Failed to write file: %v", err)
	}

	return nil
}

func renderToMarkdown(html string) (string, error) {
	now := time.Now()

	req, err := http.NewRequest("POST", "http://heckyesmarkdown.com/go/", bytes.NewBuffer([]byte("html="+url.QueryEscape(html))))
	if err != nil {
		return "", fmt.Errorf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to make request: %v", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	log.Printf("renderToMarkdown took %v", time.Since(now))

	return string(body), nil
}
