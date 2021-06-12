package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var m = make(map[interface{}]map[string][]string)
var params = make(map[string]string)
var pattern = []string{}
var version = "1.6"

// var verbose = false
var min int
var total = 0

const (
	blue  = "\u001b[38;5;77m"
	green = "\u001b[38;5;45m"
	grey  = "\u001b[38;5;252m"
	red   = "\u001b[38;5;42m"
	white = "\u001b[38;5;255m"
	reset = "\u001b[0m"
)

var banner = `                            
  ░            ░ ░      ░ ░  ░  ░            
  ░ ░        ░ ░ ░ ▒  ░ ░ ░ ▒  ░             
░░▒ ▒░    ░ ▒ ▒░   ░ ▒ ▒░ ░ ░ ▒ ░            
░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒░▒░▒░ ▒ ▒▒ ▓▒           
 ▄████▄   ▒█████   ▒█████   ██ ▄█▀           
▒██▀ ▀█  ▒██▒  ██▒▒██▒  ██▒ ██▄█▒            
▒▓█    ▄ ▒██░  ██▒▒██░  ██▒▓███▄░            
▒▓▓▄ ▄██▒▒██   ██░▒██   ██░▓██ █▄             
 ▒▓███▀ ░░ ████▓▒░░ ████▓▒░▒██▒ █▄ ` + version + `       Gitesh Sharma @giteshnxtlvl
 
 Usage  : cook [PARAM-1 VALUES] [PARAM-2 VALUES] ... [PARAM-N VALUES]  [PATTERN]
          cook -param1 val1,val2 -param2 file.txt param1:param2
 Help   : cook -h 
 Config : cook -config`

func findRegex(file, expresssion string, array *[]string) []string {
	founded := []string{}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return []string{file + ":" + expresssion}
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
		founded = append(founded, k)
	}
	return founded
}

func fileValues(file string, array *[]string) {

	readFile, err := os.Open(file)

	if err != nil {
		log.Fatalln("Err: Opening File ", file)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		*array = append(*array, fileScanner.Text())
	}
}

func appendToFile(filepath string, data []byte) {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.Write(data); err != nil {
		panic(err)
	}
}

func checkFileCache(url string, array *[]string) {
	filename := filepath.Base(url)

	err := os.MkdirAll(path.Join(home, ".cache", "cook"), os.ModePerm)
	if err != nil {
		log.Fatalln("Err: Making .cache folder in HOME/USERPROFILE ", err)
	}

	if _, err := os.Stat(path.Join(home, ".cache", "cook", filename)); err != nil {
		appendToFile(path.Join(home, ".cache", "cook", filename), getData(url))
	}

	fileValues(path.Join(home, ".cache", "cook", filename), array)
}

func updateCache() {
	fmt.Println(banner)
	for key, files := range m["files"] {
		fmt.Println("\n" + blue + key + reset)
		for _, file := range files {
			if strings.HasPrefix(file, "http://") || strings.HasPrefix(file, "https://") {
				filename := filepath.Base(file)
				// fmt.Printf("\n%s Updating %-14s:%s %s", blue, filename, reset, file)
				appendToFile(path.Join(home, ".cache", "cook", filename), getData(file))
			}
		}
	}
}

func applyCase(values []string, array *[]string, fn func(string) string) {
	for _, t := range final {
		for _, v := range values {
			*array = append(*array, t+fn(v))
		}
	}
}

func applyColumnCases(columnValues []string, columnNum int) {
	temp := []string{}

	// Using cases for columnValues
	if len(columnCases[columnNum]) > 0 {

		//All cases
		if columnCases[columnNum]["A"] {
			applyCase(columnValues, &temp, strings.ToUpper)
			applyCase(columnValues, &temp, strings.ToLower)
			applyCase(columnValues, &temp, strings.Title)
		} else {

			if columnCases[columnNum]["U"] {
				applyCase(columnValues, &temp, strings.ToUpper)
			}

			if columnCases[columnNum]["L"] {
				applyCase(columnValues, &temp, strings.ToLower)
			}

			if columnCases[columnNum]["T"] {
				applyCase(columnValues, &temp, strings.Title)
			}
		}

	} else {
		for _, t := range final {
			for _, v := range columnValues {
				temp = append(temp, t+v)
			}
		}
	}

	final = temp
}

//Initializing with empty string, so loops will run for 1st column
var final = []string{""}

func main() {
	// fmt.Fprintln(os.Stderr, banner)

	commands = os.Args[1:]
	parseInput()
	cookConfig()

	for columnNum, param := range pattern {

		columnValues := []string{}

		for _, p := range strings.Split(param, ",") {

			success := parseRanges(p, &columnValues)
			if success {
				continue
			}

			if val, exists := params[p]; exists {
				parseValue(val, &columnValues)
				continue
			}

			if val, exists := m["charSet"][p]; exists {
				columnValues = append(columnValues, strings.Split(val[0], "")...)
				continue
			}

			if files, exists := m["files"][p]; exists {

				for _, file := range files {
					if strings.HasPrefix(file, "http://") || strings.HasPrefix(file, "https://") {
						checkFileCache(file, &columnValues)
					} else {
						fileValues(file, &columnValues)
					}
				}
				continue
			}

			if val, exists := m["lists"][p]; exists {
				columnValues = append(columnValues, val...)
				continue
			}

			if val, exists := m["extensions"][p]; exists {
				for _, ext := range val {
					columnValues = append(columnValues, "."+ext)
				}
				continue
			}

			columnValues = append(columnValues, p)
		}

		applyColumnCases(columnValues, columnNum)

		if columnNum >= min {
			for _, v := range final {
				total++
				fmt.Println(v)
			}
		}
	}
}
