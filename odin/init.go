package odin

import (
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
		return
	}

	err = os.Mkdir(currentDir+"/template", 0755)
	if err != nil {
		return
	}

	err = os.Mkdir(currentDir+"/static", 0755)
	if err != nil {
		return
	}

	// create CNAME file
	emptyFile, _ := os.Create(currentDir + "/CNAME")
	err = emptyFile.Close()
	if err != nil {
		return
	}

	// copy static files
	CopyFile(rootDir+"/static/blog_list.html", currentDir+"/template/blog_list.html")
	CopyFile(rootDir+"/static/post.html", currentDir+"/template/post.html")
	CopyFile(rootDir+"/static/main.css", currentDir+"/static/main.css")

}
