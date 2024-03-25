package util

import (
	"io/ioutil"
	"log"

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

func ReadYaml(filepath string, m map[string]map[string][]string) {

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
