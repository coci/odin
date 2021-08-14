package odin

import (
	"os"
	"strings"
	"time"
)

func New(title string) {
	// replace white space with '-'
	originalTitle := title
	convertedTitle := strings.ReplaceAll(title, " ", "-")

	// get current directory
	currentDir := GetCurrentDir()

	blogPost, _ := os.Create(currentDir + "/content/"+convertedTitle+".md")

	currentTime := time.Now()
	_, err := blogPost.WriteString("---\ndate: "+currentTime.Format("2006-01-02")+"\ntitle: "+originalTitle+"\npermalink: "+convertedTitle+"\n---\n")
	if err != nil {
		return
	}

	err = blogPost.Close()
	if err != nil {
		return
	}

}
