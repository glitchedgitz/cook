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
var ConfigFolder string
var IngredientsFolder = "ingredients"

var ConfigInfo string

// Contains category and their data
var M = make(map[string]map[string][]string)
var checkM = make(map[string][]string)

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	log.Fatalln(path)
	return false
}

func firstRun() {

	fmt.Fprintln(os.Stderr, "First Run")
	fmt.Fprintln(os.Stderr, "Creating and Downloading Cook's Ingredients...\n\n ")

	err := os.MkdirAll(path.Join(ConfigFolder, IngredientsFolder), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	structure := make(map[string][]string)
	err = yaml.Unmarshal([]byte(GetData("https://raw.githubusercontent.com/glitchedgitz/cook-ingredients/main/structure")), &structure)
	if err != nil {
		log.Fatalf("Err: Parsing YAML %v", err)
	}

	for _, v := range structure["infofiles"] {
		filename := path.Base(v)
		fmt.Fprint(os.Stderr, "\rDownloading                             \r", filename)
		WriteFile(path.Join(ConfigFolder, filename), GetData(v))
	}
	for _, v := range structure["yamlfiles"] {
		filename := path.Base(v)
		fmt.Fprint(os.Stderr, "\rDownloading                             \r", filename)
		WriteFile(path.Join(ConfigFolder, IngredientsFolder, filename), GetData(v))
	}
	fmt.Fprint(os.Stderr, "\rDone                             \r")

}

func CookConfig() {
	if len(os.Getenv("COOK")) > 0 {
		ConfigFolder = os.Getenv("COOK")
	} else {
		ConfigFolder = path.Join(home, "cook-ingredients")
	}

	if !exists(ConfigFolder) {
		firstRun()
	}

	VPrint(fmt.Sprintf("Config Folder  %s", ConfigFolder))

	files, err := ioutil.ReadDir(path.Join(ConfigFolder, IngredientsFolder))
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

		ReadYaml(filename, m)

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
		ConfigInfo += fmt.Sprintf("    %-25s   %-8d %s\n", filename, total, configRows)
	}

	ConfigInfo += fmt.Sprintf("\n    %-25s   %d\n", "TOTAL FILES", totalFiles)
	ConfigInfo += fmt.Sprintf("    %-25s   %d\n", "TOTAL WORDLISTS SET", wholeTotal)

	ReadInfoYaml(path.Join(ConfigFolder, "check.yaml"), checkM)
}
