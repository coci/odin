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
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"
)

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

func SortMapByKey(m map[string][]Post) map[string][]Post {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] > keys[j] })

	finalMap := make(map[string][]Post)

	for _, k := range keys {
		finalMap[k] = m[k]
	}
	return finalMap
}

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
