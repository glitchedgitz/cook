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
		if val, exists := local[filename]; exists {
			if val[1] != "" {
				prefix = val[1] + "-"
			}
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
				M[k][prefix+kk] = vv
				total++
			}
			// M[k] = v
		}
		wholeTotal += total
		totalFiles++
		configInfo += fmt.Sprintf("    %-25s : %d\n", filename, total)
	}

	configInfo += fmt.Sprintf("\n    %-25s : %d\n", "TOTAL FILES", totalFiles)
	configInfo += fmt.Sprintf("    %-25s : %d\n", "TOTAL WORDLISTS SET", wholeTotal)
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

	fmt.Println("\nCONFIG")
	fmt.Printf("    Location: %v\n", configFolder)
	fmt.Println("\nFILES")
	fmt.Println(configInfo)

	// ShowMap("charSet")
	// ShowMap("lists")
	// ShowMap("patterns")
	// ShowMap("extensions")

	// fmt.Println("\n" + Green + strings.ToUpper("files") + Reset)

	// keys := []string{}
	// for k := range M["files"] {
	// 	keys = append(keys, k)
	// }
	// sort.Strings(keys)
	// for _, k := range keys {
	// 	files := M["files"][k]
	// 	fmt.Printf(Green+"  %-12s \n"+White, k)
	// 	for _, file := range files {

	// 		filebase := filepath.Base(file)
	// 		fmt.Println("\t" + strings.Replace(file, filebase, Green+filebase+Reset, 1))
	// 	}
	// 	fmt.Println()

	// }

	os.Exit(0)
}
