package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/glitchedgitz/cook/v2/pkg/util"
	"gopkg.in/yaml.v3"
)

func (conf *Config) getLocalFile(m map[string][]string) {
	// if len(os.Getenv("COOK")) > 0 {
	// 	ConfigFolder = os.Getenv("COOK")
	// }

	localfile := path.Join(conf.ConfigPath, "info.yaml")
	util.ReadInfoYaml(localfile, m)
}

func (conf *Config) getRepoFile(m map[string][]string) {
	content := conf.GetData("https://raw.githubusercontent.com/glitchedgitz/cook-ingredients/main/info.yaml")

	err := yaml.Unmarshal([]byte(content), &m)
	if err != nil {
		log.Fatalf("Err: Parsing YAML %s %v", "https://raw.githubusercontent.com/glitchedgitz/cook-ingredients/main/info.yaml", err)
	}
}

func (conf *Config) getConfigFiles(m map[string]bool) {
	files, err := ioutil.ReadDir(conf.ConfigPath)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		m[file.Name()] = true
	}
}

// Updating yaml file
func (conf *Config) updateFile(file string) {
	// fmt.Println("Updating : ", file)
	defer conf.wg.Done()
	content := conf.GetData("https://raw.githubusercontent.com/glitchedgitz/cook-ingredients/main/" + file)
	localFile := path.Join(conf.ConfigPath, file)
	util.WriteFile(localFile, content)
}

// Updating cook's database
func (conf *Config) UpdateDb() {
	var local = make(map[string][]string)
	var repo = make(map[string][]string)
	var files = make(map[string]bool)
	var updatedFiles = 0

	conf.getLocalFile(local)
	conf.getRepoFile(repo)
	conf.getConfigFiles(files)

	for file, values := range repo {
		version := values[0]
		if localv, exists := local[file]; exists {
			if version > localv[0] {
				conf.wg.Add(1)
				go conf.updateFile(file)
				updatedFiles++
			}
		} else if files[file] {
			log.Fatalf("\nErr: Please rename the file '%s' because cook-ingredients has new file with the same name.\n", file)
		} else {
			conf.wg.Add(1)
			fmt.Println("\nAdding new file :)")
			go conf.updateFile(file)
			updatedFiles++
		}
	}

	if updatedFiles > 0 {
		conf.wg.Add(1)
		go conf.updateFile("info.yaml")
	} else {
		fmt.Println("\nEverything is updated :)")
	}

	conf.wg.Wait()
}
