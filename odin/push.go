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
	"fmt"
	"os/exec"
)

func Push() {
	_, err := exec.Command("bash", "-c", "git add .").Output()
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
	}

	_, err = exec.Command("bash", "-c", "git commit -m 'add content'").Output()
	if err != nil {
		fmt.Println("2")
		fmt.Println(err)
	}

	_, err = exec.Command("bash", "-c", "cd blog/ && git add . && git commit -m 'publish'").Output()
	if err != nil {
		fmt.Println("3")
		fmt.Println(err)
	}

	_, err = exec.Command("bash", "-c", "cd ..").Output()
	if err != nil {
		fmt.Println("6")
		fmt.Println(err)
	}

	_, err = exec.Command("bash", "-c", "git push origin main gh-pages -ff").Output()
	if err != nil {
		fmt.Println("7")
		fmt.Println(err)
	}
}
