package cook

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"

	"gopkg.in/yaml.v3"
)

func getLocalFile(m map[string][]string) {
	if len(os.Getenv("COOK")) > 0 {
		ConfigFolder = os.Getenv("COOK")
	}

	localfile := path.Join(ConfigFolder, "info.yaml")
	ReadInfoYaml(localfile, m)
}

func getRepoFile(m map[string][]string) {
	content := GetData("https://raw.githubusercontent.com/giteshnxtlvl/cook-ingredients/main/info.yaml")

	err := yaml.Unmarshal([]byte(content), &m)
	if err != nil {
		log.Fatalf("Err: Parsing YAML %s %v", "https://raw.githubusercontent.com/giteshnxtlvl/cook-ingredients/main/info.yaml", err)
	}
}

func getConfigFiles(m map[string]bool) {
	files, err := ioutil.ReadDir(ConfigFolder)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		m[file.Name()] = true
	}
}

var wg sync.WaitGroup

// Updating yaml file
func updateFile(file string) {
	// fmt.Println("Updating : ", file)
	defer wg.Done()
	content := GetData("https://raw.githubusercontent.com/giteshnxtlvl/cook-ingredients/main/" + file)
	localFile := path.Join(ConfigFolder, file)
	WriteFile(localFile, content)
}

// Updating cook's database
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
		go updateFile("info.yaml")
	} else {
		fmt.Println("\nEverything is updated :)")
	}

	wg.Wait()
}
