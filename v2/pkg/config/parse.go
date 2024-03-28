package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/glitchedgitz/cook/v2/pkg/parse"
	"github.com/glitchedgitz/cook/v2/pkg/util"
)

func PrintFunc(k string, v []string, search string) {
	// fmt.Println(strings.ReplaceAll(k, search, "\u001b[48;5;239m"+search+Reset))
	fmt.Printf("%s%s{\n", util.Blue+k+util.Reset, strings.ReplaceAll(v[0], search, util.Blue+search+util.Reset))
	for _, file := range v[1:] {
		fmt.Printf("    %s\n", strings.ReplaceAll(file, search, util.Blue+search+util.Reset))
	}
	fmt.Print("}\n\n")
}

// Checking for functions
func (conf *Config) ParseFunc(value string, array *[]string) bool {

	if !(strings.Contains(value, "[") && strings.Contains(value, "]")) {
		return false
	}

	funcName, funcArgs := parse.ReadSqBrSepBy(value, ",")
	// fmt.Println(funcName)
	// fmt.Println(funcArgs)

	fmt.Print("")

	if funcPatterns, exists := conf.Ingredients["functions"][funcName]; exists {

		funcDef := strings.Split(funcPatterns[0][1:len(funcPatterns[0])-1], ",")

		// fmt.Printf("Func Arg: %v", funcArgs)
		// fmt.Printf("\tFunc Def: %v", funcDef)

		if len(funcDef) != len(funcArgs) {
			log.Fatalln("\nErr: No of Arguments are different for")
			PrintFunc(funcName, funcPatterns, funcName)
		}

		for _, p := range funcPatterns[1:] {
			for index, arg := range funcDef {
				p = strings.ReplaceAll(p, arg, funcArgs[index])
			}
			*array = append(*array, p)
		}

		return true
	}
	return false
}

func (conf *Config) ParseFile(param string, value string, array *[]string) bool {

	if conf.InputFile[param] {
		if strings.Contains(value, "https://") || strings.Contains(value, "http://") {
			// Parse URL
			URLValues(value, array)
		} else {
			// When you paste URL from chrome
			// file:///C://Users//gites//AppData//Local/grroxy/downloaded_lL5A
			value = strings.TrimLeft(value, "file:///")
			FileValues(value, array)
		}
		return true
	}

	if conf.checkFileSet(value, array) {
		return true
	}

	return false
}

func PipeInput(value string, array *[]string) bool {
	var pipe []string
	if value == "-" {
		sc := bufio.NewScanner(os.Stdin)
		if len(pipe) > 0 {
			*array = append(*array, pipe...)
		}
		for sc.Scan() {
			*array = append(*array, sc.Text())
			pipe = append(pipe, sc.Text())
		}
		return true
	}
	return false
}

func RawInput(value string, array *[]string) bool {
	if value == "`" {
		return true
	}
	if strings.HasPrefix(value, "`") && strings.HasSuffix(value, "`") {
		lv := len(value)
		*array = append(*array, []string{value[1 : lv-1]}...)
		return true
	}
	return false
}

func ParseRanges(p string, array *[]string) bool {

	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	if strings.Count(p, "-") == 1 {

		numRange := strings.SplitN(p, "-", 2)
		from := numRange[0]
		to := numRange[1]

		start, err1 := strconv.Atoi(from)
		stop, err2 := strconv.Atoi(to)

		if err1 == nil && err2 == nil {
			for start <= stop {
				*array = append(*array, strconv.Itoa(start))
				start++
			}
			return true
		}

		if len(from) == 1 && len(to) == 1 && strings.Contains(chars, from) && strings.Contains(chars, to) {
			start = strings.Index(chars, from)
			stop = strings.Index(chars, to)

			if start < stop {
				charsList := strings.Split(chars, "")
				for start <= stop {
					*array = append(*array, charsList[start])
					start++
				}
				return true
			}
		}
	}
	return false
}

func ParsePorts(ports []string, array *[]string) {

	for _, p := range ports {
		if ParseRanges(p, array) {
			continue
		}
		port, err := strconv.Atoi(p)
		if err != nil {
			log.Printf("Err: Is this port number -_-?? '%s'", p)
		}
		*array = append(*array, strconv.Itoa(port))
	}
}
