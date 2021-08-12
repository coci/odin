package odin

import (
	"fmt"
	"log"
	"os"
)

func getCurrentDir() string {
	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}
	return path
}
func Init() {
	currentDir := getCurrentDir()
	fmt.Println(currentDir)
}
