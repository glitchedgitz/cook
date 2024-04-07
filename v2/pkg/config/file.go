package config

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/glitchedgitz/cook/v2/pkg/util"
)

func AddFilesToArray(value string, array *[]string) {
	tmp := make(map[string]bool)
	RawFileValues(value, tmp)
	for k := range tmp {
		*array = append(*array, k)
	}
}

func FileValues(pattern string, array *[]string, peek int) {
	files, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatalln("Err: In pattern ", err)
	}
	for _, file := range files {
		ReadFile, err := os.Open(file)

		if err != nil {
			log.Fatalln("Err: Opening File ", file)
		}

		defer ReadFile.Close()

		fileScanner := bufio.NewScanner(ReadFile)

		for fileScanner.Scan() {
			line := fileScanner.Text()
			line = strings.TrimRight(line, "\r")
			*array = append(*array, line)
			if peek > 0 {
				peek--
				if peek == 0 {
					break
				}
			}
		}

		if err := fileScanner.Err(); err != nil {
			log.Fatal(err)
		}

	}
}

func RawFileValues(pattern string, allLines map[string]bool) {
	files, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatalln("Err: In pattern ", err)
	}
	for _, file := range files {
		ReadFile, err := os.Open(file)

		if err != nil {
			log.Fatalln("Err: Opening File ", file)
		}

		defer ReadFile.Close()

		fileScanner := bufio.NewScanner(ReadFile)

		for fileScanner.Scan() {
			line := fileScanner.Text()
			line = strings.TrimRight(line, "\n")
			line = strings.TrimRight(line, "\r")
			if allLines[line] {
				continue
			}
			allLines[line] = true
		}
		if err := fileScanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func FindRegex(data []byte, expresssion string, array *[]string) {

	r, err := regexp.Compile(expresssion)
	if err != nil {
		log.Fatalln(err)
	}

	e := make(map[string]bool)

	// replacing \r (carriage return) as it puts cursor on beginning of line
	for _, found := range r.FindAllString(strings.ReplaceAll(string(data), "\r", ""), -1) {
		if e[found] {
			continue
		}
		e[found] = true
		*array = append(*array, found)
	}
}

func (conf *Config) FileRegex(file string, expresssion string, array *[]string) {
	FindRegex(util.ReadFile(file), expresssion, array)
}

func (conf *Config) RawFileRegex(files []string, expresssion string, array *[]string) {
	for _, file := range files {
		if strings.HasPrefix(file, "http://") || strings.HasPrefix(file, "https://") {
			FindRegex(conf.GetData(file), expresssion, array)
		} else {
			conf.FileRegex(file, expresssion, array)
		}
	}
}
