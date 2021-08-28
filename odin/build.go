// Copyright (C) 2021- Soroush Safari <soroush_safarii@yahoo.com>
//
// This file is part of Odin.
//
// Odin is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Odin is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.
//
//You should have received a copy of the GNU General Public License
// along with Odin.  If not, see <http://www.gnu.org/licenses/>.

package odin

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// IndexPage : we use this struct to ship data into index.html
type IndexPage struct {
	Title    string            // title of owner of blog ( like : soroush safari )
	BlogPost map[string][]Post // all posts
}

// create blog directory
func clearBlogDir() {
	currentDir := GetCurrentDir()

	d, err := os.Open(currentDir + "/blog")
	if err != nil {
		log.Println(err)
	}

	defer d.Close()

	names, _ := d.Readdirnames(-1)

	for _, name := range names {
		if filepath.Join(currentDir+"/blog", name) != currentDir+"/blog/.git" {
			_ = os.RemoveAll(filepath.Join(currentDir+"/blog", name))
		}
	}
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

	_ = os.Mkdir(currentDir+"/blog/static", 0755)

	CopyFile(currentDir+"/static/highlight.pack.js", currentDir+"/blog/static/highlight.pack.js")
	CopyFile(currentDir+"/static/main.css", currentDir+"/blog/static/main.css")

	CopyFile(currentDir+"/CNAME", currentDir+"/blog/CNAME")
}

// list all post in content directory
func listPosts() []string {
	var postList []string

	currentDirectory := GetCurrentDir()

	files, err := ioutil.ReadDir(currentDirectory + "/content")
	if err != nil {
		log.Println(err)
	}

	for _, element := range files {
		if element.Name() != ".DS_Store" {
			postList = append(postList, element.Name())
		}
	}
	return postList
}

// read Metadata from head of markdown file
// in the Metadata there is ( date , title , permalink) which is useful data
func readMeta(source []byte) (time.Time, string, string) {
	// markdown extension that read Meta from head of markdown file
	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	// read Meta data
	var buf bytes.Buffer
	context := parser.NewContext()
	if err := markdown.Convert(source, &buf, parser.WithContext(context)); err != nil {
		log.Println(err)
	}
	metaData := meta.Get(context)

	// get Meta data
	date, _ := time.Parse("2006-01-02", metaData["date"].(string))

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

	// find size of metadata
	rs := re.FindStringSubmatch(buf.String())

	// return post content exclude meta data part
	return buf.String()[len(rs[0]):]

}

// create index.html that list all blog posts
func buildIndex(posts []Post) {
	currentDir := GetCurrentDir()

	context := IndexPage{}
	context.Title = "soroush"

	var postMap = make(map[string][]Post)

	for _, e := range posts {
		postMap[strconv.Itoa(e.Date.Year())] = append(postMap[strconv.Itoa(e.Date.Year())], e)
	}

	context.BlogPost = SortMapByKey(postMap)

	// the original template
	originIndexHtmlTemplate, _ := ioutil.ReadFile(currentDir + "/template/index.html")

	var templateBuffer bytes.Buffer
	// concat context into original template
	tpl := template.Must(template.New("originIndexHtmlTemplate").Parse(string(originIndexHtmlTemplate)))
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
	cfg := ReadConfig()
	// convert title to 'dash-seperated'
	convertedTitle := strings.ReplaceAll(post.Title, " ", "-")

	// get content of post exclude Metadata part
	content := readContent(post)

	// create directory for post
	createDirForPost(currentDir, convertedTitle)

	// read per-exist template
	postHtmlTemplate, _ := ioutil.ReadFile(currentDir + "/template/post.html")

	// replace content
	tpl := template.Must(template.New("postHtmlTemplate").Parse(string(postHtmlTemplate)))
	var templateBuffer bytes.Buffer

	// convert post date to string
	year, month, day := post.Date.Date()
	stringDate := fmt.Sprintf("%d-%d-%d", year, month, day)

	if cfg.Site.Language == "fa" {
		yearFa, monthFa, dayFa := ptime.New(post.Date).Date()
		stringDate = fmt.Sprintf("%d-%d-%d", yearFa, monthFa, dayFa)
	}

	context := map[string]string{
		"Title":     post.Title,
		"Date":      stringDate,
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

	// clear blog dir from old files ( previous build )
	clearBlogDir()

	// copy required file
	copyRequiredFile()

	// list all posts
	postList := listPosts()

	// store all permalinks for posts
	var posts []Post

	for _, dir := range postList {
		currentDir := GetCurrentDir()
		// read post
		source, _ := ioutil.ReadFile(currentDir + "/content/" + dir)

		// get Metadata from the post
		date, title, permalink := readMeta(source)

		// create slug
		splitDate := strings.Split(date.Format("2006-01-02"), "-")
		slugifyDate := strings.Join(splitDate[1:], "/")

		// create Post type
		blogPost := Post{Date: date, Slug: slugifyDate, Title: title, Permalink: permalink, Source: source}
		posts = append(posts, blogPost)
		// make html from post
		buildPost(&blogPost)
	}

	// make index.html that list all blog posts
	buildIndex(posts)
}
