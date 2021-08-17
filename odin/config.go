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
