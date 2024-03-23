package config

import (
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func ReadFile(filepath string) []byte {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalln("Err: Reading File ", filepath, err)
	}
	return content
}

func WriteFile(filepath string, data []byte) {
	err := ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		log.Fatalln("Err: Writing File ", filepath, err)
	}
}

func ReadYaml(filename string, m map[string]map[string][]string) {
	filepath := path.Join(ConfigFolder, IngredientsFolder, filename)

	content := ReadFile(filepath)

	err := yaml.Unmarshal([]byte(content), &m)
	if err != nil {
		log.Fatalf("Err: Parsing YAML %s %v", filepath, err)
	}
}

func WriteYaml(filepath string, m interface{}) {
	data, err := yaml.Marshal(&m)

	if err != nil {
		log.Fatal(err)
	}

	WriteFile(filepath, data)
}

func ReadInfoYaml(filepath string, m map[string][]string) {
	content := ReadFile(filepath)

	err := yaml.Unmarshal([]byte(content), &m)
	if err != nil {
		log.Fatalf("Err: Parsing YAML %s %v", filepath, err)
	}
}

func checkFileSet(p string, array *[]string) bool {
	if files, exists := M["files"][p]; exists {

		CheckFileCache(p, files)
		FileValues(path.Join(home, ".cache", "cook", p), array)
		return true

	} else if files, exists := M["raw-files"][p]; exists {

		tmp := make(map[string]bool)
		for _, file := range files {
			if strings.HasPrefix(file, "http://") || strings.HasPrefix(file, "https://") {
				RawFileValues(path.Join(home, ".cache", "cook", filepath.Base(file)), tmp)
			} else {
				RawFileValues(file, tmp)
			}
		}

		for k := range tmp {
			*array = append(*array, k)
		}
		return true

	}

	return false
}

func CheckYaml(p string, array *[]string) bool {

	if val, exists := M["lists"][p]; exists {
		*array = append(*array, val...)
		return true
	}

	if val, exists := M["ports"][p]; exists {
		ParsePorts(val, array)
		return true
	}

	if checkFileSet(p, array) {
		return true
	}

	return false
}
