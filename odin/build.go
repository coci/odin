package odin

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// create dir for each post in /blog dir
func createDirForPost(currentDir, title string) {
	_, err := os.Stat(currentDir+"/blog/"+title)

	// check if there isn't dir
	if os.IsNotExist(err) {
		err = os.Mkdir(currentDir+"/blog/"+title, 0755)
		if err != nil {
			log.Println(err)
		}
	}
}

// Post struct will contain blog post
type Post struct {
	date      string
	title     string
	permalink string // link of post
	source    []byte // content of markdown file
}

// copy required static file
func copyStatic() {
	currentDir := GetCurrentDir()

	err := os.Mkdir(currentDir+"/blog/static", 0755)
	if err != nil {
		return
	}

	CopyFile(currentDir+"/static/highlight.pack.js", currentDir+"/blog/static/highlight.pack.js")
	CopyFile(currentDir+"/static/main.css", currentDir+"/blog/static/main.css")
}

// read Meta data from head of markdown file
func readMeta(source []byte) (string, string, string) {
	// extension for reading Meta data from head of markdown file
	// in the Meta data there is ( date , title , permalink) which is useful data

	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	var buf bytes.Buffer
	context := parser.NewContext()
	if err := markdown.Convert(source, &buf, parser.WithContext(context)); err != nil {
		log.Panicln()
	}
	metaData := meta.Get(context)

	// get Meta data
	date := metaData["date"].(string)
	title := metaData["title"].(string)
	permalink := metaData["permalink"].(string)

	return date, title, permalink
}

// read data from markdown exclude Meta data
func readContent(post *Post) string {
	var buf bytes.Buffer
	if err := goldmark.Convert(post.source, &buf); err != nil {
		panic(err)
	}

	// slice of content of post exclude Meta data part
	return buf.String()[69:]
}

// create index.html that list all blog posts
func buildIndex(postList []string) {
	//
}

// create html file from markdown file
func buildPost(post *Post) {
	currentDir := GetCurrentDir()
	projectRoot := GetProjectRootDir()

	// convert title to 'dash-seperated'
	convertedTitle := strings.ReplaceAll(post.title, " ", "-")

	// get content of post exclude Meta data part
	content := readContent(post)

	// create directory for post
	createDirForPost(currentDir, convertedTitle)

	// read per-exist template
	postHtmlTemplate, _ := ioutil.ReadFile(projectRoot + "/static/post.html")

	// replace content
	tpl := template.Must(template.New("postHtmlTemplate").Parse(string(postHtmlTemplate)))
	var templateBuffer bytes.Buffer
	context := map[string]string{
		"Title":     post.title,
		"Date":      post.date,
		"Content":   content,
		"Permalink": post.permalink,
	}

	_ = tpl.Execute(&templateBuffer, context)

	// create html file and write pre-exists post template
	htmlFile, _ := os.Create(currentDir + "/blog/" + convertedTitle + "/index.html")

	// write (post from markdown and concat the post into per-existing tempalte ) into html file
	_, _ = htmlFile.WriteString(templateBuffer.String())

	err := htmlFile.Close()
	if err != nil {
		panic(err)
	}

}

// Build is entry function for 'odin build' command
func Build() {
	// copy required static file
	copyStatic()

	// list all posts
	postList := ListPosts()

	for _, post := range postList {
		currentDir := GetCurrentDir()
		// read post
		source, _ := ioutil.ReadFile(currentDir + "/content/" + post)

		// get Meta data from the post
		date, title, permalink := readMeta(source)

		// create Post type
		post := Post{date: date, title: title, permalink: permalink, source: source}

		// make html from post
		buildPost(&post)
	}

	// make index.html that list all blog posts
	buildIndex(postList)
}
