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
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

// SortBlogPostByYear sort blog posts
func SortBlogPostByYear(posts []Post) []map[string][]Post {
	// this function will return something like this :
	// [ map{"2021":[post,post,post,....} , map{"2020":[post,post,post,....} , ..... ]
	// the key in map are year of post
	// the order of map added to slice is in reverse order

	// simulation of set data structure
	// yearsMap contain unique key of  year of posts
	var yearsMap = make(map[string]string)

	for _, e := range posts {
		_, ok := yearsMap[strings.Split(e.Date, "-")[0]]
		if !ok {
			yearsMap[strings.Split(e.Date, "-")[0]] = strings.Split(e.Date, "-")[0]
		}
	}

	// create slice of unique years
	var yearList []string
	for k := range yearsMap {
		yearList = append(yearList, k)
	}

	// reverse sort of years
	// this will be something like : ["2021" "2020" "2019"]
	sort.Slice(yearList, func(i, j int) bool { return yearList[i] > yearList[j] })

	// blogPosts contain all posts with key as year and value as post in reverse order of years like :
	// [ map{"2021":[post,post,post,....} , map{"2020":[post,post,post,....} , ..... ]
	var blogPost []map[string][]Post
	for _, e := range yearList {
		var yearPost = make(map[string][]Post)
		for _, p := range posts {
			if strings.Split(p.Date, "-")[0] == e {
				yearPost[e] = append(yearPost[e], p)
			}
		}
		blogPost = append(blogPost, yearPost)

	}

	return blogPost
}

// ReadConfig read config.yaml
func ReadConfig() OdinConfig {
	var cfg OdinConfig
	currentDir := GetCurrentDir()

	f, _ := os.Open(currentDir + "/config.yaml")
	configData := yaml.NewDecoder(f)
	err := configData.Decode(&cfg)
	if err != nil {
		log.Println(err)
	}

	return cfg
}

func GetCurrentDir() string {
	// find current directory of user terminal
	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}
	return path
}

func CopyFile(src, dst string) {
	// copy file from src to dst

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		log.Println(err)
	}

	if !sourceFileStat.Mode().IsRegular() {
		log.Println(src)
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
