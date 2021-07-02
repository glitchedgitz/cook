package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

var content []byte
var home, _ = os.UserHomeDir()
var configFolder = `E:\tools\base\cook`
var configInfo string
var M = make(map[string]map[string][]string)

func CookConfig() {

	if len(os.Getenv("COOK")) > 0 {
		configFolder = os.Getenv("COOK")
	}

	VPrint(fmt.Sprintf("Config Folder  %s", configFolder))

	files, err := ioutil.ReadDir(configFolder)
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
		if !strings.HasSuffix(filename, ".yaml") {
			continue
		}

		prefix := ""
		configRows := ""
		if val, exists := local[filename]; exists {

			v, p, r := val[0], val[1], val[2]
			if p != "" {
				prefix = p + "-"
			}

			configRows = fmt.Sprintf("%-4s   %-6s   %s", v, p, r)
		}

		content, err = ioutil.ReadFile(path.Join(configFolder, filename))
		if err != nil {
			log.Fatalln("Err: Reading Config File ", err)
		}

		err := yaml.Unmarshal([]byte(content), &m)
		if err != nil {
			log.Fatalf("Err: Parsing YAML %s %v", filename, err)
		}

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
	fmt.Println(Green + "\n    CONFIG" + Reset)
	fmt.Printf("    Location: %v\n", configFolder)
	fmt.Printf(Green+"\n    %-25s   %s     %s   %s   %s\n"+Reset, "FILE", "SETS", "VERN", "PREFIX", "REPO")
	fmt.Println(configInfo)

	os.Exit(0)
}
