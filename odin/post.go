package odin

import "time"

// Post : this struct will contain blog post
type Post struct {
	Date      time.Time
	Slug      string // like : month / day ( 01/26 ) . used in index.html page for list all blog posts
	Title     string // post title
	Permalink string // link of post
	Source    []byte // content of markdown file
}
