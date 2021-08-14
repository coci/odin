package odin

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type Post struct {
	date      string
	title     string
	permalink string
	source    []byte
}

func readMeta(source []byte) (string, string, string) {

	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)
	var buf bytes.Buffer
	context := parser.NewContext()
	if err := markdown.Convert(source, &buf, parser.WithContext(context)); err != nil {
		panic(err)
	}
	metaData := meta.Get(context)

	date := metaData["date"].(string)
	title := metaData["title"].(string)
	permalink := metaData["permalink"].(string)

	return date, title, permalink
}

func readContent(post *Post) string {
	var buf bytes.Buffer
	if err := goldmark.Convert(post.source, &buf); err != nil {
		panic(err)
	}

	return buf.String()[69:]
}

func buildIndex(postList []string) {
	//
}

func buildPost(post *Post) {
	currentDir := GetCurrentDir()
	projectRoot := GetProjectRootDir()

	// convert title to 'dash-seperated'
	convertedTitle := strings.ReplaceAll(post.title, " ", "-")

	content := readContent(post)

	// create
	err := os.Mkdir(currentDir+"/blog/"+convertedTitle, 0755)
	if err != nil {
		return
	}

	// read and create post with post pre-exists template
	postHtmlTemplate, _ := ioutil.ReadFile(projectRoot + "/static/post.html")

	// replace content
	template := template.Must(template.New("postHtmlTemplate").Parse(string(postHtmlTemplate)))
	var templateBuffer bytes.Buffer
	context := map[string]string{
		"Title":post.title,
		"Date":post.date,
		"Content":content,
	}
	template.Execute(&templateBuffer,context)


	// create html file and write pre-exists post template
	htmlFile, _ := os.Create(currentDir + "/blog/" + convertedTitle + "/index.html")

	// write final data into html file
	htmlFile.WriteString(templateBuffer.String())

	err = htmlFile.Close()
	if err != nil {
		return
	}

}
func Build() {
	// list of all posts
	//var posts []string

	postList := ListPosts()

	buildIndex(postList)

	for _, post := range postList {
		currentDir := GetCurrentDir()
		source, _ := ioutil.ReadFile(currentDir + "/content/" + post)
		date, title, permalink := readMeta(source)

		post := Post{date: date, title: title, permalink: permalink, source: source}
		buildPost(&post)
	}
}
