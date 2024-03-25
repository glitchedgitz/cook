package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/glitchedgitz/cook/v2/pkg/util"
	"gopkg.in/yaml.v3"
)

// First Run
func (conf *Config) FirstRun() {

	fmt.Fprintln(os.Stderr, "First Run")
	fmt.Fprintln(os.Stderr, "Creating and Downloading Cook's Ingredients...\n\n ")

	err := os.MkdirAll(conf.ConfigPath, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.MkdirAll(conf.IngredientsPath, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.MkdirAll(conf.CachePath, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	structure := make(map[string][]string)
	err = yaml.Unmarshal([]byte(conf.GetData("https://raw.githubusercontent.com/glitchedgitz/cook-ingredients/main/structure")), &structure)
	if err != nil {
		log.Fatalf("Err: Parsing YAML %v", err)
	}

	for _, v := range structure["infofiles"] {
		filename := path.Base(v)
		fmt.Fprint(os.Stderr, "\rDownloading                             \r", filename)
		util.WriteFile(path.Join(conf.ConfigPath, filename), conf.GetData(v))
	}
	for _, v := range structure["yamlfiles"] {
		filename := path.Base(v)
		fmt.Fprint(os.Stderr, "\rDownloading                             \r", filename)
		util.WriteFile(path.Join(conf.IngredientsPath, filename), conf.GetData(v))
	}
	fmt.Fprint(os.Stderr, "\rDone                             \r")

}

func (conf *Config) CookConfig() {

	if conf.ReConfigure || !util.Exists(conf.ConfigPath) {
		conf.FirstRun()
	}

	conf.VPrint(fmt.Sprintf("Config Folder  %s", conf.ConfigPath))

	files, err := ioutil.ReadDir(conf.IngredientsPath)
	if err != nil {
		log.Fatalln(err)
	}

	wholeTotal := 0
	totalFiles := 0

	var local = make(map[string][]string)
	conf.getLocalFile(local)

	conf.M = make(map[string]map[string][]string)
	conf.CheckM = make(map[string][]string)

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

		util.ReadYaml(path.Join(conf.IngredientsPath, filename), m)

		total := 0
		for k, v := range m {
			if _, exists := conf.M[k]; !exists {
				conf.M[k] = make(map[string][]string)
			}

			for kk, vv := range v {
				conf.M[k][prefix+strings.ToLower(kk)] = vv
				total++
			}
		}
		wholeTotal += total
		totalFiles++
		// Temporary Commented
		conf.ConfigInfo += fmt.Sprintf("    %-25s   %-8d %s\n", filename, total, configRows)
	}

	conf.ConfigInfo += fmt.Sprintf("\n    %-25s   %d\n", "TOTAL FILES", totalFiles)
	conf.ConfigInfo += fmt.Sprintf("    %-25s   %d\n", "TOTAL WORDLISTS SET", wholeTotal)

	util.ReadInfoYaml(path.Join(conf.ConfigPath, "check.yaml"), conf.CheckM)
}
