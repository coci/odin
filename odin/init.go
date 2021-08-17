package odin

import (
	"log"
	"os"
)

func Init() {
	//get current directory
	currentDir := GetCurrentDir()

	// get project root directory
	rootDir := GetProjectRootDir()

	// create necessary directories
	err := os.Mkdir(currentDir+"/content", 0755)
	if err != nil {
		log.Println(err)
	}

	err = os.Mkdir(currentDir+"/template", 0755)
	if err != nil {
		log.Println(err)
	}

	err = os.Mkdir(currentDir+"/static", 0755)
	if err != nil {
		log.Println(err)
	}

	// create CNAME file
	emptyFile, _ := os.Create(currentDir + "/CNAME")
	err = emptyFile.Close()
	if err != nil {
		log.Println(err)
	}

	// copy static files
	CopyFile(rootDir+"/static/index.html", currentDir+"/template/index.html")
	CopyFile(rootDir+"/static/post.html", currentDir+"/template/post.html")
	CopyFile(rootDir+"/static/main.css", currentDir+"/static/main.css")
	CopyFile(rootDir+"/static/highlight.pack.js", currentDir+"/static/highlight.pack.js")

}
