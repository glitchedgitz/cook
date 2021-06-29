package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"sync"

	"gopkg.in/yaml.v3"
)

func getLocalFile(m map[string][]string) {
	if len(os.Getenv("COOK")) > 0 {
		configFolder = os.Getenv("COOK")
	}

	localfile := path.Join(configFolder, "info")

	content, err := ioutil.ReadFile(localfile)

	if err != nil {
		log.Fatalf("Err: Reading File %s \n%v", localfile, err)
	}

	err = yaml.Unmarshal([]byte(content), &m)
	if err != nil {
		log.Fatalf("Err: Parsing YAML %s %v", localfile, err)
	}
}

func getRepoFile(m map[string][]string) {
	content := GetData("https://raw.githubusercontent.com/giteshnxtlvl/cooks-wordlists-database/main/info")

	err := yaml.Unmarshal([]byte(content), &m)
	if err != nil {
		log.Fatalf("Err: Parsing YAML %s %v", "https://raw.githubusercontent.com/giteshnxtlvl/cooks-wordlists-database/main/info", err)
	}
}

func getConfigFiles(m map[string]bool) {
	files, err := ioutil.ReadDir(configFolder)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		m[file.Name()] = true
	}
}

var wg sync.WaitGroup

func updateFile(file string) {
	// fmt.Println("Updating : ", file)
	defer wg.Done()
	content := GetData("https://raw.githubusercontent.com/giteshnxtlvl/cooks-wordlists-database/main/" + file)
	localFile := path.Join(configFolder, file)
	err := ioutil.WriteFile(localFile, content, 0644)
	if err != nil {
		panic(err)
	}
}

func PrintConfig(m map[string][]string) {
	fmt.Printf("\n%-25s   %s   %s   %s\n\n", "FILE", "VERSION", "PREFIX", "REPO")
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%-25s   %-7s   %-6s   %s\n", k, m[k][0], m[k][1], m[k][2])
	}
}

func UpdateCook() {
	var local = make(map[string][]string)
	var repo = make(map[string][]string)
	var files = make(map[string]bool)
	var updatedFiles = 0

	getLocalFile(local)
	getRepoFile(repo)
	getConfigFiles(files)

	for file, values := range repo {
		version := values[0]
		if localv, exists := local[file]; exists {
			if version > localv[0] {
				wg.Add(1)
				go updateFile(file)
				updatedFiles++
			}
		} else if files[file] {
			log.Fatalf("\nErr: Please rename the file '%s' because cook-wordlist-database has new file with the same name.\n", file)
		} else {
			wg.Add(1)
			fmt.Println("\nAdding new file :)")
			go updateFile(file)
			updatedFiles++

		}
	}

	if updatedFiles > 0 {
		wg.Add(1)
		go updateFile("info")
	} else {
		fmt.Println("\nEverything is updated :)")
	}

	wg.Wait()
}
