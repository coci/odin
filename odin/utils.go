package odin

import (
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func GetCurrentDir() string {
	// find current directory of user terminal
	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}
	return path
}

func GetProjectRootDir() string {
	// get odin installation dir
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func CopyFile(src, dst string) {
	// copy file from src to dst

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
			log.Println(err)
		}
	}(source)

	destination, err := os.Create(dst)
	if err != nil {
		log.Println(err)
	}
	defer func(destination *os.File) {
		err := destination.Close()
		if err != nil {
			log.Println(err)
		}
	}(destination)
	_, err = io.Copy(destination, source)
	if err != nil {
		log.Println(err)
	}

}