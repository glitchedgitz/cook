package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

var content []byte
var home, _ = os.UserHomeDir()
var configFile = path.Join(home, ".config", "cook", "cook.yaml")
var M = make(map[interface{}]map[string][]string)

func CookConfig() {

	if len(ConfigPath) > 0 {
		configFile = ConfigPath
	} else if len(os.Getenv("COOK")) > 0 {
		configFile = os.Getenv("COOK")
	}

	VPrint(fmt.Sprintf("Config File  %s", configFile))

	if _, err := os.Stat(configFile); err == nil {

		content, err = ioutil.ReadFile(configFile)
		if err != nil {
			log.Fatalln("Err: Reading Config File ", err)
		}

		if len(content) == 0 {
			fmt.Println("Downloading/Updating cook.yaml...")

			config := GetData("https://raw.githubusercontent.com/giteshnxtlvl/cook/main/cook.yaml")
			ioutil.WriteFile(configFile, []byte(config), 0644)
			content = []byte(config)
		}

	} else {

		err := os.MkdirAll(path.Join(home, ".config", "cook"), os.ModePerm)
		if err != nil {
			log.Fatalln("Err: Making .config folder in HOME/USERPROFILE ", err)
		}

		fmt.Println("Downloading/Updating cook.yaml...")

		config := GetData("https://raw.githubusercontent.com/giteshnxtlvl/cook/main/cook.yaml")
		err = ioutil.WriteFile(configFile, []byte(config), 0644)
		if err != nil {
			log.Fatalln("Err: Writing Config File", err)
		}
		content = []byte(config)
	}

	err := yaml.Unmarshal([]byte(content), &M)

	if err != nil {
		log.Fatalln("Err: Parsing YAML", err)
	}
}

func ShowMap(set string) {
	fmt.Println("\n" + Green + strings.ToUpper(set) + Reset)

	keys := []string{}
	for k := range M[set] {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf(Blue+"  %-12s "+White+"%v\n", k, M[set][k])
	}
}

func ShowConfig() {

	fmt.Println(Green + "\nCOOK.YAML " + Reset)
	fmt.Printf(Blue+"  %-11s "+White+" %v\n", "Location", configFile)

	ShowMap("charSet")
	ShowMap("files")
	ShowMap("lists")
	ShowMap("patterns")
	ShowMap("extensions")

	os.Exit(0)
}
