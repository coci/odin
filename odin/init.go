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
