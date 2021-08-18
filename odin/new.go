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
	"log"
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
		log.Println(err)
	}

	err = blogPost.Close()
	if err != nil {
		log.Println(err)
	}

}
