package core

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

func GetData(url string) []byte {
	fmt.Fprintf(os.Stderr, "Fetching: %s\n", url)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()
	return data
}

func CheckFileCache(url string) {
	filename := filepath.Base(url)

	err := os.MkdirAll(path.Join(home, ".cache", "cook"), os.ModePerm)
	if err != nil {
		log.Fatalln("Err: Making .cache folder in HOME/USERPROFILE ", err)
	}

	if _, err := os.Stat(path.Join(home, ".cache", "cook", filename)); err != nil {
		AppendToFile(path.Join(home, ".cache", "cook", filename), GetData(url))
	}
}

func UpdateCache() {
	fmt.Println(Banner)

	goaddresses := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		go func() {
			for file := range goaddresses {
				func(file string) {
					defer wg.Done()
					filename := filepath.Base(file)
					filepath := path.Join(home, ".cache", "cook", filename)
					os.Remove(filepath)
					AppendToFile(filepath, GetData(file))
				}(file)
			}
		}()
	}

	for _, files := range M["files"] {
		// fmt.Println("\n" + Blue + key + Reset)

		for _, file := range files {
			if strings.HasPrefix(file, "http://") || strings.HasPrefix(file, "https://") {
				wg.Add(1)
				goaddresses <- file
			}
		}
	}

	wg.Wait()
}

func AppendToFile(filepath string, data []byte) {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.Write(data); err != nil {
		panic(err)
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
