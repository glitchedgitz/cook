package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"cook/pkg/core"
)

// var version = "1.6"

var total = 0

// var home, _ = os.UserHomeDir()

var columnCases = make(map[int]map[string]bool)

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

	params, pattern := core.ParseInput()

	var home, _ = os.UserHomeDir()
	core.CookConfig()

	for columnNum, param := range pattern {

		columnValues := []string{}
		var success bool

		for _, p := range strings.Split(param, ",") {

			success = core.ParseRanges(p, &columnValues)
			if success {
				continue
			}

			if val, exists := params[p]; exists {
				core.ParseValue(val, &columnValues)
				continue
			}

			// Raw String using `
			if strings.HasPrefix(p, "`") && strings.HasSuffix(p, "`") {
				lv := len(p)
				columnValues = append(columnValues, []string{p[1 : lv-1]}...)
				continue
			}

			if val, exists := core.M["charSet"][p]; exists {
				columnValues = append(columnValues, strings.Split(val[0], "")...)
				continue
			}

			if files, exists := core.M["files"][p]; exists {
				for _, file := range files {
					if strings.HasPrefix(file, "http://") || strings.HasPrefix(file, "https://") {
						core.CheckFileCache(file)
						core.FileValues(path.Join(home, ".cache", "cook", filepath.Base(file)), &columnValues)
					} else {
						core.FileValues(file, &columnValues)
					}
				}
				continue
			}

			if val, exists := core.M["lists"][p]; exists {
				columnValues = append(columnValues, val...)
				continue
			}

			if val, exists := core.M["extensions"][p]; exists {
				for _, ext := range val {
					columnValues = append(columnValues, "."+ext)
				}
				continue
			}

			columnValues = append(columnValues, p)
		}

		applyColumnCases(columnValues, columnNum)

		if columnNum >= core.Min {
			for _, v := range final {
				total++
				fmt.Println(v)
			}
		}
	}

	core.VPrint(fmt.Sprintf("Total words generated: %d", total))
}
