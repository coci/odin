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

}
