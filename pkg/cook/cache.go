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

	if _, err := os.Stat(filepath); err != nil {
		var tmp = make(map[string]bool)
		f, err := os.OpenFile(filepath, os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		for _, file := range files {
			VPrint(fmt.Sprintf("GetData(): Fetching %s\n", file))

			res, err := http.Get(file)
			if err != nil {
				log.Fatal(err)
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
					log.Fatal(err)
				}
			}

			if err := fileScanner.Err(); err != nil {
				log.Fatal(err)
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

func AppendToFile(filepath string, data []byte) {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err = f.Write(data); err != nil {
		log.Fatal(err)
	}
}

func UpdateCache() {

	type filedata struct {
		filepath string
		files    []string
	}

	goaddresses := make(chan filedata)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go func() {
			for f := range goaddresses {
				CheckFileCache(f.filepath, f.files)
				wg.Done()
			}
		}()
	}

	for k, files := range checkM {
		wg.Add(1)
		goaddresses <- filedata{path.Join(home, ".cache", "cook", k), files}
	}

	wg.Wait()
}
