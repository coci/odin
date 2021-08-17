package odin

import (
	"bytes"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
)

// Post struct will contain blog post
type Post struct {
	Date      string
	Title     string
	Permalink string // link of post
	Source    []byte // content of markdown file
}

type IndexPage struct {
	Title    string
	BlogPost []Post
}

// create dir for each post in /blog dir
func createDirForPost(currentDir, title string) {
	_, err := os.Stat(currentDir + "/blog/" + title)

	// check if there isn't dir
	if os.IsNotExist(err) {
		err = os.Mkdir(currentDir+"/blog/"+title, 0755)
		if err != nil {
			log.Println(err)
		}
	}
}

// copy required file ( static files, CNAME)
func copyRequiredFile() {
	currentDir := GetCurrentDir()

	err := os.Mkdir(currentDir+"/blog/static", 0755)
	if err != nil {
		log.Println(err)
	}

	CopyFile(currentDir+"/static/highlight.pack.js", currentDir+"/blog/static/highlight.pack.js")
	CopyFile(currentDir+"/static/main.css", currentDir+"/blog/static/main.css")

	CopyFile(currentDir+"/CNAME", currentDir+"/blog/CNAME")
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
		log.Println(err)
	}
	metaData := meta.Get(context)

	// get Meta data
	date := metaData["date"].(string)
	title := metaData["title"].(string)
	permalink := metaData["permalink"].(string)

	return date, title, permalink
}

// read data from markdown exclude Meta data and convert it to html
func readContent(post *Post) string {
	var buf bytes.Buffer
	var re = regexp.MustCompile(`<hr>(.|\n)*?</h2>`)

	// convert markdown string to html
	if err := goldmark.Convert(post.Source, &buf); err != nil {
		log.Println(err)
	}

	// slice of content of post exclude Meta data part

	rs:=re.FindStringSubmatch(buf.String())

	return buf.String()[len(rs[0]):]

}

// create index.html that list all blog posts
func buildIndex(posts []Post) {
	currentDir := GetCurrentDir()
	var templateBuffer bytes.Buffer

	context := IndexPage{
		"test",
		posts,
	}

	IndexHtmlTemplate, _ := ioutil.ReadFile(currentDir + "/template/index.html")
	tpl := template.Must(template.New("postHtmlTemplate").Parse(string(IndexHtmlTemplate)))
	_ = tpl.Execute(&templateBuffer, context)

	// create html file and write pre-exists post template
	htmlFile, _ := os.Create(currentDir + "/blog/index.html")

	// write into html file
	_, _ = htmlFile.WriteString(templateBuffer.String())

	err := htmlFile.Close()
	if err != nil {
		log.Println(err)
	}

}

// create html file from markdown file
func buildPost(post *Post) {
	currentDir := GetCurrentDir()
	projectRoot := GetProjectRootDir()

	// convert title to 'dash-seperated'
	convertedTitle := strings.ReplaceAll(post.Title, " ", "-")

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
		"Title":     post.Title,
		"Date":      post.Date,
		"Content":   content,
		"Permalink": post.Permalink,
	}

	_ = tpl.Execute(&templateBuffer, context)

	// create html file and write pre-exists post template
	htmlFile, _ := os.Create(currentDir + "/blog/" + convertedTitle + "/index.html")

	// write (post from markdown and concat the post into per-existing tempalte ) into html file
	_, _ = htmlFile.WriteString(templateBuffer.String())

	err := htmlFile.Close()
	if err != nil {
		log.Println(err)
	}

}

// Build is entry function for 'odin build' command
func Build() {
	// copy required file
	copyRequiredFile()

	// list all posts
	postList := ListPosts()

	// store all permalinks for posts
	var posts []Post

	for _, dir := range postList {
		currentDir := GetCurrentDir()
		// read post
		source, _ := ioutil.ReadFile(currentDir + "/content/" + dir)

		// get Meta data from the post
		date, title, permalink := readMeta(source)

		// create Post type
		blogPost := Post{Date: date, Title: title, Permalink: permalink, Source: source}
		posts = append(posts, blogPost)
		// make html from post
		buildPost(&blogPost)

	}

	// make index.html that list all blog posts
	buildIndex(posts)
}
