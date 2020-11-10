package main

import (
	"flag"
	"io/ioutil"
	"strings"
)

var Blue = "\033[96m"
var White = "\033[97m"

func commonfiles() {

}

func main() {
	var words string
	flag.StringVar(&words, "w", "", "help message for flagname")

	var extensions string
	flag.StringVar(&extensions, "e", "", "help message for flagname")

	var file string
	flag.StringVar(&file, "f", "", "help message for flagname")
	flag.Parse()

	if file != "" {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		println(string(data))
	}

	extensionArray := strings.Split(extensions, ",")
	wordsArray := strings.Split(words, ",")

	for _, word := range wordsArray {
		for _, ext := range extensionArray {
			println(word + ext)
		}
	}

}
