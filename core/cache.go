package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

func GetData(url string) []byte {
	VPrint(fmt.Sprintf("Fetching: %s\n", url))

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

func UpdateCache() {
	fmt.Println(Banner)

	goaddresses := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		go func() {
			for file := range goaddresses {
				func(file string) {
					defer wg.Done()
					fmt.Println("Updating ", file)
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
