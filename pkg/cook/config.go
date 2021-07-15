package cook

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

// var content []byte
var home, _ = os.UserHomeDir()
var ConfigFolder = `E:\tools\base\cook`
var configInfo string
var M = make(map[string]map[string][]string)

var checkM = make(map[string][]string)

func readCheckYaml() {
	filepath := path.Join(ConfigFolder, "check.yaml")
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalln("Err: Reading File ", filepath, err)
	}

	err = yaml.Unmarshal([]byte(content), &checkM)
	if err != nil {
		log.Fatalf("Err: Parsing YAML %s %v", filepath, err)
	}
}

func ReadYaml(filepath string, m map[string]map[string][]string) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalln("Err: Reading File ", filepath, err)
	}

	err = yaml.Unmarshal([]byte(content), &m)
	if err != nil {
		log.Fatalf("Err: Parsing YAML %s %v", filepath, err)
	}
}

func writeYaml(filepath string, m map[string][]string) {
	data, err := yaml.Marshal(&m)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filepath, data, 0)
	if err != nil {
		log.Fatalln("Err: Writing File ", filepath, err)
	}
}

func CookConfig() {

	if len(os.Getenv("COOK")) > 0 {
		ConfigFolder = os.Getenv("COOK")
	}

	VPrint(fmt.Sprintf("Config Folder  %s", ConfigFolder))

	filesFolders := path.Join(ConfigFolder, "yaml")

	files, err := ioutil.ReadDir(filesFolders)
	if err != nil {
		log.Fatalln(err)
	}

	wholeTotal := 0
	totalFiles := 0

	var local = make(map[string][]string)
	getLocalFile(local)

	for _, file := range files {
		var m = make(map[string]map[string][]string)

		filename := file.Name()

		prefix := ""
		configRows := ""
		if val, exists := local[filename]; exists {

			v, p, r := val[0], val[1], val[2]
			if p != "" {
				prefix = p + "-"
			}

			configRows = fmt.Sprintf("%-4s   %-6s   %s", v, p, r)
		}

		ReadYaml(path.Join(filesFolders, filename), m)

		total := 0
		for k, v := range m {
			if _, exists := M[k]; !exists {
				M[k] = make(map[string][]string)
			}

			for kk, vv := range v {
				M[k][prefix+strings.ToLower(kk)] = vv
				total++
			}
		}
		wholeTotal += total
		totalFiles++
		configInfo += fmt.Sprintf("    %-25s   %-8d %s\n", filename, total, configRows)
	}

	configInfo += fmt.Sprintf("\n    %-25s   %d\n", "TOTAL FILES", totalFiles)
	configInfo += fmt.Sprintf("    %-25s   %d\n", "TOTAL WORDLISTS SET", wholeTotal)

	readCheckYaml()
}

func ShowConfig() {
	fmt.Println(Blue + "\n    CONFIG" + Reset)
	fmt.Printf("    Location: %v\n", ConfigFolder)
	fmt.Printf(Blue+"\n    %-25s   %s     %s   %s   %s\n"+Reset, "FILE", "SETS", "VERN", "PREFIX", "REPO")
	fmt.Println(configInfo)

	os.Exit(0)
}
