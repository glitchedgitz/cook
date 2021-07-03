package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/giteshnxtlvl/cook/core"
	"github.com/giteshnxtlvl/cook/parse"
)

func leetBegin() {
	leetValues["0"] = []string{"o", "O"}
	leetValues["1"] = []string{"i", "I", "l", "L"}
	leetValues["3"] = []string{"e", "E"}
	leetValues["4"] = []string{"a", "A"}
	leetValues["5"] = []string{"s", "S"}
	leetValues["6"] = []string{"b"}
	leetValues["7"] = []string{"t", "T"}
	leetValues["8"] = []string{"B"}
}

func analyseParams(params map[string]string) {
	for param, value := range params {
		// fmt.Println(params)
		if strings.HasSuffix(param, ":") {
			delete(params, param)
			param = strings.TrimSuffix(param, ":")
			core.InputFile[param] = true
			params[param] = value
		}
	}
}

func searchMode(cmds []string) {
	core.CookConfig()

	search := cmds[0]
	found := false

	for cat, vv := range core.M {
		for k, v := range vv {
			k = strings.ToLower(k)

			if strings.Contains(k, search) {
				fmt.Println()
				if cat == "files" {
					fmt.Println(strings.ReplaceAll(k, search, "\u001b[48;5;239m"+search+core.Reset))
					for _, file := range v {
						fmt.Printf("    %s\n", strings.ReplaceAll(file, search, core.Green+search+core.Reset))
					}

				} else {
					fmt.Printf("%s \n\t%v\n", k, v)
				}
				found = true
			}
		}
	}

	if !found {
		fmt.Println("Not Found: ", search)
	}
	os.Exit(0)
}

func addMode(cmds []string) {
}
func updateMode(cmds []string) {
}
func deleteMode(cmds []string) {
}

func parseInput() (map[string]string, []string) {

	parse.Help = core.Banner
	parse.Parse()

	if help {
		core.ShowHelp()
	}

	if showConfig {
		core.CookConfig()
		core.ShowConfig()
	}

	if update != "" {
		if update == "cook" {
			core.UpdateCook()
		}
		// core.CookConfig()
		// core.UpdateCache()
		os.Exit(0)
	}

	if len(encodeValue) > 0 {
		encodeString = strings.Split(encodeValue, ",")
		finalFunc = encode
	}

	core.Verbose = verbose

	params = parse.UserDefinedFlags()
	analyseParams(params)

	pattern := parse.Args
	noOfColumns := len(pattern)

	if noOfColumns > 0 {
		if pattern[0] == "search" {
			searchMode(pattern[1:])
		} else if pattern[0] == "help" {
			core.HelpMode(pattern[1:])
		} else if pattern[0] == "add" {
			addMode(pattern[1:])
		} else if pattern[0] == "update" {
			updateMode(pattern[1:])
		} else if pattern[0] == "delete" {
			deleteMode(pattern[1:])
		}
	}

	if Min < 0 {

		Min = noOfColumns - 1
	} else {
		if Min > noOfColumns {
			fmt.Println("Err: Min is greator than no of columns")
			os.Exit(0)
		}
		Min -= 1
	}

	if caseValue != "" {
		columnCases = core.UpdateCases(caseValue, noOfColumns)
	}

	if l337 > -1 {
		if l337 > 1 {
			fmt.Println("Err: -1337 can be 0 or 1, 0 - Calm Mode & 1 - Angry Mode", l337)
			os.Exit(0)
		}
		leetBegin()
	}

	if len(appendColumns) > 0 {
		columns := strings.Split(appendColumns, ",")
		for _, colNum := range columns {
			intValue, err := strconv.Atoi(colNum)
			if err != nil {
				log.Fatalf("Err: Column Value %s in not integer", colNum)
			}
			appendMode[intValue] = true
		}
	}

	if showCol {
		fmt.Fprintln(os.Stderr)
		for i, p := range pattern {
			fmt.Fprintf(os.Stderr, "Col %d: %s\n", i, p)
		}
		fmt.Fprintln(os.Stderr)
		os.Exit(0)
	}

	core.VPrint(fmt.Sprintf("Pattern: %v \n", pattern))

	return params, pattern
}
