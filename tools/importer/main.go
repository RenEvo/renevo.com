package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync/atomic"
	"time"

	"github.com/metal3d/go-slugify"
	"github.com/mmcdole/gofeed"
	"github.com/op/go-logging"
)

var minWordCount = flag.Int("min-words", 100, "The minimum amount of words to keep")

func main() {
	flag.Parse()
	logger := logging.MustGetLogger("importer")
	var loggerFormat = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} ▶ %{level:.4s} %{color:reset} %{message}`,
	)
	logging.SetFormatter(loggerFormat)

	stdOutBackend := logging.NewLogBackend(os.Stdout, "", 0)

	// Only errors and more severe messages should be sent to backend1
	levelLogger := logging.AddModuleLevel(stdOutBackend)
	levelLogger.SetLevel(logging.WARNING, "")

	// Set the backends to be used.
	logger.SetBackend(levelLogger)

	blogFeeds := []string{
		"http://renevo.com/blogs/developer/atom.aspx",
		"http://renevo.com/blogs/community_blogs/atom.aspx",
		"http://renevo.com/blogs/dotnet/atom.aspx",
	}

	blogImporter := NewBlogImporter(logger, &BlogImporterConfig{"./imported", blogFeeds, *minWordCount})

	blogImporter.Parse()

	logger.Infof("Finished Downloading %v items", blogImporter.PostCount())
}

type BlogImporterConfig struct {
	OutputDir    string
	Feeds        []string
	MinWordCount int
}

type BlogImporter struct {
	logger    *logging.Logger
	parser    *gofeed.Parser
	postCount uint32
	config    *BlogImporterConfig
}

func NewBlogImporter(logger *logging.Logger, config *BlogImporterConfig) *BlogImporter {
	config.OutputDir = path.Clean(config.OutputDir)

	return &BlogImporter{
		logger: logger,
		parser: gofeed.NewParser(),
		config: config,
	}
}

func (this *BlogImporter) Parse() {
	now := time.Now()

	for _, feed := range this.config.Feeds {
		this.parseFeed(feed)
		this.logger.Infof("Finished: %s", feed)
	}

	this.logger.Infof("Parse took: %v", time.Since(now))
}

func (this *BlogImporter) PostCount() uint32 {
	return this.postCount
}

func (this *BlogImporter) parseFeed(feed string) {
	now := time.Now()

	parsed, err := this.parser.ParseURL(feed)

	if err != nil {
		this.logger.Fatalf("Failed to parse feed: %v", err)
	}

	this.logger.Infof(parsed.Title)

	feedUrl, _ := url.Parse(feed)
	feedBasePath, _ := path.Split(feedUrl.Path)
	completion := make(chan *gofeed.Item)
	defer close(completion)

	for _, item := range parsed.Items {
		go func(thisItem *gofeed.Item) {
			now := time.Now()
			if err := this.parseFeedItem(feedBasePath, thisItem); err != nil {
				this.logger.Error(err.Error())
			} else {
				atomic.AddUint32(&this.postCount, 1)
			}
			this.logger.Debugf("Parse Feed Item %s took: %v", thisItem.Title, time.Since(now))
			completion <- thisItem
		}(item)
	}

	for i := 0; i < len(parsed.Items); i++ {
		<-completion
	}

	this.logger.Debugf("Parse Feed %s took: %v", feed, time.Since(now))
}
func (this *BlogImporter) parseFeedItem(feedBasePath string, item *gofeed.Item) error {
	if !strings.EqualFold(item.Author.Name, "tom anderson") {
		this.logger.Warningf("Skipping %v's post %s", item.Author.Name, item.Title)
		return nil
	}

	postUrl, _ := url.Parse(item.Link)
	postPath, postFileName := path.Split(postUrl.Path)

	dir := path.Join(this.config.OutputDir, feedBasePath, fmt.Sprintf("%d", item.PublishedParsed.Year()), fmt.Sprintf("%02d", item.PublishedParsed.Month()), fmt.Sprintf("%02d", item.PublishedParsed.Day()))
	os.MkdirAll(dir, 0644)

	slug := slugify.Marshal(item.Title, true)
	filePrefix := path.Join(dir, slug)

	// save html version
	this.logger.Debugf("Post: '%v' by '%v' posted on '%v'", item.Title, item.Author.Name, item.PublishedParsed)
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

	// detect and download images
	postImages := getImageList(bodyString)
	if len(postImages) > 0 {
		for _, image := range postImages {
			if strings.Contains(image.url, "aggbug.aspx?PostID=") {
				bodyString = removeImage(image, bodyString)
				continue // skipping these things
			}
			os.MkdirAll(filePrefix, 0644)
			this.logger.Debugf("Found Image [%s]: %s @ url: %s", image.id, image.title, image.url)
			outputPath, err := downloadImage(image, filePrefix)
			if err != nil {
				this.logger.Errorf("Post: '%s': '%s' failed to download image: %v", item.Title, item.Link, err)
				continue
			}
			if len(outputPath) == 0 {
				// image doesn't exist, prune it from existance
				bodyString = removeImage(image, bodyString)
			} else {
				bodyString = updateImage(image, dir, outputPath, bodyString)
				// image exists, update the markdown to point to the relative dir
				this.logger.Infof("Image: %s", outputPath)
			}
		}
	}

	if err := ioutil.WriteFile(filePrefix+".md", append([]byte(yamlFrontMatter), []byte(bodyString)...), 0644); err != nil {
		return fmt.Errorf("ERR: Failed to write file: %v", err)
	}

	return nil
}

func downloadImage(src *PostImage, dest string) (string, error) {
	req, err := http.NewRequest("GET", src.url, nil)
	if err != nil {
		return "", fmt.Errorf("Failed to create request: %v", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to make request: %v", err)
	}
	if res.StatusCode == http.StatusNotFound {
		return "", nil
	} else if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Server returned a not good error code: [%v] %v", res.StatusCode, src.url)
	}
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	contentType := res.Header.Get("Content-Type")
	fileExtension, err := mime.ExtensionsByType(contentType)
	if err != nil {
		return "", err
	}
	outputFileName := dest
	if uri, err := url.Parse(src.url); err != nil {
		return "", fmt.Errorf("Failed to find a suitable name for image %s", src.id)
	} else {

		outputFileName = path.Join(dest, strings.TrimSuffix(path.Base(uri.Path), path.Ext(uri.Path)))
	}

	for _, ext := range fileExtension {
		if err := ioutil.WriteFile(outputFileName+ext, body, 0644); err != nil {
			return "", fmt.Errorf("ERR: Failed to write file: %v", err)
		}

		return outputFileName + ext, nil
	}

	return "", fmt.Errorf("No suitable file extensions for content type: %s", contentType)
}

func renderToMarkdown(html string) (string, error) {
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

	return cleanMarkdown(string(body)), nil
}

var doubleLineBreaks = regexp.MustCompile(`\n\s\s\s\s\n`)
var codeStart = regexp.MustCompile(`\s+1:`)
var hrFix = regexp.MustCompile(`\n\*\s\*\s\*\n`)

func cleanMarkdown(markdown string) string {
	result := doubleLineBreaks.ReplaceAllString(markdown, "")
	result = codeStart.ReplaceAllString(result, fmt.Sprint("\n\n       1:"))
	result = hrFix.ReplaceAllString(result, fmt.Sprint("\n***\n"))
	return result
}

var imageTagRegex = regexp.MustCompile(`!\[([^\]]*)\]\[(\d+)\]`)
var linkRefRegex = regexp.MustCompile(`\[(\d+)\]:\s([^\n].*)`)

type PostImage struct {
	id    string
	title string
	url   string
}

func updateImage(src *PostImage, baseDir string, newSrc string, markdown string) string {
	relative, _ := filepath.Rel(baseDir, newSrc)
	relative = "./" + strings.Replace(relative, "\\", "/", -1)
	return strings.Replace(markdown, src.url, relative, -1)
}

func removeImage(src *PostImage, markdown string) string {
	tagRegex := regexp.MustCompile(`!\[([^\]]*)\]\[(` + src.id + `)\]`)
	linkRegex := regexp.MustCompile(`\[(` + src.id + `)\]:\s([^\n].*)`)

	markdown = tagRegex.ReplaceAllString(markdown, "")
	markdown = linkRegex.ReplaceAllString(markdown, "")

	return markdown
}

func getImageList(markdown string) []*PostImage {
	images := []*PostImage{}
	imageMap := make(map[string]*PostImage)

	for _, match := range imageTagRegex.FindAllStringSubmatch(markdown, -1) {
		postImage := &PostImage{id: match[2], title: match[1]}
		images = append(images, postImage)
		imageMap[match[2]] = postImage
	}

	for _, match := range linkRefRegex.FindAllStringSubmatch(markdown, -1) {
		if postImage, ok := imageMap[match[1]]; ok {
			postImage.url = match[2]
			// some weird thing that happens somewhere, this fixes it
			if strings.HasSuffix(postImage.url, ` "`+postImage.title+`"`) {
				postImage.url = strings.Replace(postImage.url, ` "`+postImage.title+`"`, "", -1)
			}
		}
	}

	return images
}
