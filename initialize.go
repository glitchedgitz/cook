package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/giteshnxtlvl/cook/core"
	"github.com/giteshnxtlvl/cook/parse"
)

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
				} else if cat == "raw-files" {
					fmt.Println(strings.ReplaceAll(k, search, "\u001b[48;5;239m"+search+core.Reset))
					for _, file := range v {
						fmt.Printf("    %s\n", strings.ReplaceAll(file, search, core.Green+search+core.Reset))
					}
				} else if cat == "patterns" {
					fmt.Println(strings.ReplaceAll(k, search, "\u001b[48;5;239m"+search+core.Reset))
					fmt.Printf("    %s%s{\n", k, strings.ReplaceAll(v[0], search, core.Green+search+core.Reset))
					for _, file := range v[1:] {
						fmt.Printf("\t%s\n", strings.ReplaceAll(file, search, core.Green+search+core.Reset))
					}
					fmt.Println("    }")
				} else {
					fmt.Println(strings.ReplaceAll(k, search, "\u001b[48;5;239m"+search+core.Reset))
					fmt.Println(strings.ReplaceAll(fmt.Sprintf("    %v\n", v), search, core.Green+search+core.Reset))
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

var home, _ = os.UserHomeDir()

func updateMode(cmds []string) {
	filename := cmds[0]
	filepath := path.Join(home, ".cache", "cook", filename)
	if files, exists := core.M["files"][filename]; exists {
		os.Remove(filepath)
		core.CheckFileCache(filename, files)
	}
}

func deleteMode(cmds []string) {
}
func cleanMode(cmds []string) {
}
func infoMode(cmds []string) {
}

var cmdFunctions = map[string]func([]string){
	"search": searchMode,
	"help":   core.HelpMode,
	"add":    addMode,
	"clean":  cleanMode,
	"info":   infoMode,
	"update": updateMode,
	"delete": deleteMode,
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
		doEncode = true
	}

	core.Verbose = verbose

	params = parse.UserDefinedFlags()
	analyseParams(params)

	pattern := parse.Args
	noOfColumns := len(pattern)

	if noOfColumns > 0 {
		if fn, exists := cmdFunctions[pattern[0]]; exists {
			fn(pattern[1:])
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
		doLeet = true
		if l337 > 1 {
			fmt.Println("Err: -1337 can be 0 or 1, 0 - Calm Mode & 1 - Angry Mode", l337)
			os.Exit(0)
		}
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
