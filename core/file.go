package core

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func FileValues(file string, array *[]string) {

	readFile, err := os.Open(file)

	if err != nil {
		log.Fatalln("Err: Opening File ", file)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		line := strings.TrimRight(fileScanner.Text(), "\r")
		*array = append(*array, line)
	}
}

func FindRegex(files []string, expresssion string, array *[]string) {

	for _, file := range files {
		if strings.HasPrefix(file, "http://") || strings.HasPrefix(file, "https://") {
			CheckFileCache(file)
			file = path.Join(home, ".cache", "cook", filepath.Base(file))
		}

		content, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalln("Error reading file" + file)
		}

		r, err := regexp.Compile(expresssion)
		if err != nil {
			log.Fatalln(err)
		}

		e := make(map[string]bool)
		// replacing \r (carriage return) as cursor moves to start of the line
		for _, found := range r.FindAllString(strings.ReplaceAll(string(content), "\r", ""), -1) {
			e[found] = true
		}

		for k := range e {
			*array = append(*array, k)
		}
	}
}
