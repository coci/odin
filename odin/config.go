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
	"log"
	"os/exec"
)

func Config() {
	var repo string

	fmt.Println("please enter github repository url :")
	_, err := fmt.Scan(&repo)
	if err != nil {
		return
	};

	_, err = exec.Command("bash", "-c", "git init").Output()
	if err != nil {
		log.Println(err)
	}

	_, err = exec.Command("bash", "-c", "echo \"blog\" > .gitignore").Output()
	if err != nil {
		log.Println(err)
	}

	_, err = exec.Command("bash", "-c", "echo \".DS_Store\" > .gitignore").Output()
	if err != nil {
		log.Println(err)
	}

	_, err = exec.Command("bash", "-c", "git add --all").Output()
	if err != nil {
		log.Println(err)
	}

	_, err = exec.Command("bash", "-c", "git commit -m 'initial project'").Output()
	if err != nil {
		log.Println(err)
	}

	_, err = exec.Command("bash", "-c", "git branch -M main").Output()
	if err != nil {
		log.Println(err)
	}

	_, err = exec.Command("bash", "-c", "git remote add origin "+repo).Output()
	if err != nil {
		log.Println(err)
	}

	_, err = exec.Command("bash", "-c", "git push -u origin main").Output()
	if err != nil {
		log.Println(err)
	}

	_, err = exec.Command("bash", "-c", "git checkout --orphan gh-pages").Output()
	if err != nil {
		log.Println(err)
	}

	_, err = exec.Command("bash", "-c", "git reset --hard").Output()
	if err != nil {
		log.Println(err)
	}

	_, err = exec.Command("bash", "-c", "git commit --allow-empty -m \"Initializing gh-pages\"").Output()
	if err != nil {
		log.Println(err)
	}

	_, err = exec.Command("bash", "-c", "git push origin gh-pages").Output()
	if err != nil {
		log.Println(err)
	}

	_, err = exec.Command("bash", "-c", "git checkout main").Output()
	if err != nil {
		log.Println(err)
	}

	_, err = exec.Command("bash", "-c", "git worktree add -B gh-pages blog origin/gh-pages").Output()
	if err != nil {
		log.Println(err)
	}

}
