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

	_, err = exec.Command("bash", "-c", "git push origin main gh-pages").Output()
	if err != nil {
		fmt.Println("7")
		fmt.Println(err)
	}
}