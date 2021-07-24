package cook

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"sync"
)

// Return data from url
func GetData(url string) []byte {
	VPrint(fmt.Sprintf("GetData(): Fetching %s\n", url))

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()
	return data
}

func makeCacheFolder() {
	err := os.MkdirAll(path.Join(home, ".cache", "cook"), os.ModePerm)
	if err != nil {
		log.Fatalln("Err: Making .cache folder in HOME/USERPROFILE ", err)
	}
}

// Checking if file's cache present
func CheckFileCache(filename string, files []string) {

	makeCacheFolder()
	filepath := path.Join(home, ".cache", "cook", filename)

	if _, e := os.Stat(filepath); e != nil {
		fmt.Fprintf(os.Stderr, "Creating cache for %s\n", filename)
		var tmp = make(map[string]bool)
		f, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("Creating File: ", err)
		}

		defer f.Close()

		for _, file := range files {
			fmt.Fprintf(os.Stderr, "Fetching %s\n", file)

			res, err := http.Get(file)

			if err != nil {
				log.Fatal("Getting Data", err)
			}

			defer res.Body.Close()

			fileScanner := bufio.NewScanner(res.Body)

			for fileScanner.Scan() {
				line := fileScanner.Text()
				if tmp[line] {
					continue
				}
				tmp[line] = true
				if _, err = f.WriteString(fileScanner.Text() + "\n"); err != nil {
					log.Fatalf("Writing File: %v", err)
				}
			}

			if err := fileScanner.Err(); err != nil {
				log.Fatalf("FileScanner: %v", err)
			}
		}
		checkM[filename] = files
		WriteYaml(path.Join(ConfigFolder, "check.yaml"), checkM)

	} else {

		chkfiles := checkM[filename]
		if len(files) != len(chkfiles) {
			os.Remove(filepath)
			CheckFileCache(filename, files)
			WriteYaml(path.Join(ConfigFolder, "check.yaml"), checkM)
			return
		}
		for i, v := range chkfiles {
			if v != files[i] {
				os.Remove(filepath)
				CheckFileCache(filename, files)
				WriteYaml(path.Join(ConfigFolder, "check.yaml"), checkM)
				break
			}
		}
	}
}

func UpdateCache() {

	type filedata struct {
		filename string
		files    []string
	}

	goaddresses := make(chan filedata)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go func() {
			for f := range goaddresses {
				filepath := path.Join(home, ".cache", "cook", f.filename)
				os.Remove(filepath)
				CheckFileCache(f.filename, f.files)
				wg.Done()
			}
		}()
	}

	for filename, files := range checkM {
		wg.Add(1)
		goaddresses <- filedata{filename, files}
	}

	wg.Wait()
	fmt.Fprintf(os.Stderr, "All files updated")
}
