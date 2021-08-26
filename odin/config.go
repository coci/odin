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
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os/exec"
)

type OdinConfig struct {
	Site struct {
		Author   string `yaml:"author"`
		Domain   string `yaml:"domain"`
		Language string `yaml:"language"`
		Github   string `yaml:"github"`
	} `yaml:"site"`
}

func configGit(repo string) {
	_, err := exec.Command("bash", "-c", "git init").Output()
	if err != nil {
		log.Println(err)
	}

	_, err = exec.Command("bash", "-c", "echo '.DS_Store\nblog' > .gitignore").Output()
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

	_, err = exec.Command("bash", "-c", "git push -u origin main -ff").Output()
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

	_, err = exec.Command("bash", "-c", "git push origin gh-pages -ff").Output()
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

func Config() {
	var repo, language, author string

	currentDir := GetCurrentDir()

	// create config.yaml in blog repository
	err := ioutil.WriteFile(currentDir+"/config.yaml", []byte(ConfigYaml), 0644)
	if err != nil {
		log.Println(err)
	}

	// config object
	cfg := ReadConfig()

	// get github repo
	fmt.Println("please enter github repository url :")
	_, err = fmt.Scan(&repo)
	if err != nil {
		log.Println(err)
	}

	// add setting
	cfg.Site.Github = repo

	// get site language
	fmt.Println("please enter blog language ( fa/en ) :")
	_, err = fmt.Scan(&language)
	if err != nil {
		log.Println(err)
	}

	// add setting
	cfg.Site.Language = language

	// get author
	fmt.Println("please enter your name ( it will show in title of blog ):")
	_, err = fmt.Scan(&author)
	if err != nil {
		log.Println(err)
	}

	// add setting
	cfg.Site.Author = author

	// marshal config obj with new data
	changedData, _ := yaml.Marshal(cfg)

	// write new config in config.yaml
	err = ioutil.WriteFile(currentDir+"/config.yaml", changedData, 0644)
	if err != nil {
		log.Println(err)
	}


	// copy static files
	if cfg.Site.Language == "en" {
		err := ioutil.WriteFile(currentDir+"/template/index.html", []byte(IndexEn), 0644)
		if err != nil {
			log.Println(err)
		}
		err = ioutil.WriteFile(currentDir+"/template/post.html", []byte(PostEn), 0644)
		if err != nil {
			log.Println(err)
		}
		err = ioutil.WriteFile(currentDir+"/static/main.css", []byte(Css), 0644)
		if err != nil {
			log.Println(err)
		}
		err = ioutil.WriteFile(currentDir+"/static/highlight.pack.js", []byte(Js), 0644)
		if err != nil {
			log.Println(err)
		}

	} else if cfg.Site.Language == "fa" {
		err = ioutil.WriteFile(currentDir+"/template/index.html", []byte(IndexFa), 0644)
		if err != nil {
			log.Println(err)
		}
		err = ioutil.WriteFile(currentDir+"/template/post.html", []byte(PostFa), 0644)
		if err != nil {
			log.Println(err)
		}
		err = ioutil.WriteFile(currentDir+"/static/main.css", []byte(Css), 0644)
		if err != nil {
			log.Println(err)
		}
		err = ioutil.WriteFile(currentDir+"/static/highlight.pack.js", []byte(Js), 0644)
		if err != nil {
			log.Println(err)
		}
	}

	// configure git
	configGit(repo)

}
