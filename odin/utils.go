package odin

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func GetCurrentDir() string {
	// find current directory
	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}
	return path
}

func GetProjectRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func CopyFile(src, dst string) {
	// copy files through project

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		log.Println(err)
	}

	if !sourceFileStat.Mode().IsRegular() {
		log.Println("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		log.Println(err)
	}
	defer func(source *os.File) {
		err := source.Close()
		if err != nil {

		}
	}(source)

	destination, err := os.Create(dst)
	if err != nil {
		log.Println(err)
	}
	defer func(destination *os.File) {
		err := destination.Close()
		if err != nil {

		}
	}(destination)
	io.Copy(destination, source)

}


func ListPosts() []string {
	var postList []string

	currentDirectory := GetCurrentDir()

	files, err := ioutil.ReadDir(currentDirectory +"/content")
	if err != nil {
		log.Println(err)
	}

	for _ ,element := range files{
		postList = append(postList,element.Name())
	}
	return postList
}