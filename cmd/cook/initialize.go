package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/giteshnxtlvl/cook/pkg/cook"
	"github.com/giteshnxtlvl/cook/pkg/parse"
)

func analyseParams(params map[string]string) {
	for param, value := range params {
		// fmt.Println(params)
		if strings.HasSuffix(param, ":") {
			delete(params, param)
			param = strings.TrimSuffix(param, ":")
			cook.InputFile[param] = true
			params[param] = value
		}
	}
}

func searchMode(cmds []string) {

	search := cmds[0]
	found := false

	for cat, vv := range cook.M {
		for k, v := range vv {
			k = strings.ToLower(k)

			if strings.Contains(k, search) {
				fmt.Println()
				if cat == "files" || cat == "raw-files" {
					fmt.Println(strings.ReplaceAll(k, search, "\u001b[48;5;239m"+search+cook.Reset))
					for _, file := range v {
						fmt.Printf("    %s\n", strings.ReplaceAll(file, search, cook.Blue+search+cook.Reset))
					}
				} else if cat == "patterns" {
					cook.PrintPattern(k, v, search)
				} else {
					fmt.Println(strings.ReplaceAll(k, search, "\u001b[48;5;239m"+search+cook.Reset))
					fmt.Println(strings.ReplaceAll(fmt.Sprintf("    %v\n", v), search, cook.Blue+search+cook.Reset))
				}
				found = true
			}
		}
	}

	if !found {
		fmt.Println("Not Found: ", search)
	}

}

func addMode(cmds []string) {
}

var home, _ = os.UserHomeDir()

func updateMode(cmds []string) {
	filename := cmds[0]
	filepath := path.Join(home, ".cache", "cook", filename)
	if files, exists := cook.M["files"][filename]; exists {
		os.Remove(filepath)
		cook.CheckFileCache(filename, files)
	}
}

func deleteMode(cmds []string) {
}
func cleanMode(cmds []string) {
}
func infoMode(cmds []string) {
	set := cmds[0]

	filepath := path.Join(cook.ConfigFolder, "yaml", set)

	m := make(map[string]map[string][]string)
	if strings.HasSuffix(set, ".yaml") || strings.HasPrefix(set, ".yml") {
		cook.ReadYaml(filepath, m)
	}

	fmt.Println("\n" + cook.Blue + set + cook.Reset)
	fmt.Println("Path    : ", filepath)
	fmt.Println("Sets    : ", len(m))
	fmt.Println("Version : ", len(m))
}

func showMode(cmds []string) {
	set := cmds[0]

	if strings.HasSuffix(set, ".yaml") || strings.HasPrefix(set, ".yml") {
		data, err := ioutil.ReadFile(path.Join(cook.ConfigFolder, "yaml", set))
		fmt.Println()
		fmt.Println(string(data))
		if err == nil {
			return
		}
	}

	if vals, exists := cook.M[set]; exists {
		fmt.Printf("\n" + cook.Blue + strings.ToUpper(set) + cook.Reset + "\n\n")

		keys := []string{}
		for k := range vals {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		if set != "patterns" {
			for _, k := range keys {
				fmt.Printf("  %-12s "+"%v\n", k, vals[k])
			}
		} else {
			for _, k := range keys {
				cook.PrintPattern(k, vals[k], "")
			}
		}
	} else {
		fmt.Println("\nNot Found " + set + "\nTry charset, extensions, patterns, files, raw-files, ports or [file.yaml")
	}
}

var cmdFunctions = map[string]func([]string){
	"search": searchMode,
	"show":   showMode,
	"help":   cook.HelpMode,
	"add":    addMode,
	"clean":  cleanMode,
	"info":   infoMode,
	"update": updateMode,
	"delete": deleteMode,
	"size":   cook.TerminalSize,
}

func parseInput() (map[string]string, []string) {
	parse.Help = cook.Banner
	parse.Parse()

	if help {
		cook.ShowHelp()
	}

	if showConfig {
		cook.CookConfig()
		cook.ShowConfig()
	}

	if update != "" {
		if update == "cook" {
			cook.UpdateCook()
		}
		// cook.CookConfig()
		// cook.UpdateCache()
		os.Exit(0)
	}

	if len(encodeValue) > 0 {
		encodeString = strings.Split(encodeValue, ",")
		doEncode = true
	}

	cook.Verbose = verbose

	params = parse.UserDefinedFlags()
	analyseParams(params)

	pattern := parse.Args
	noOfColumns := len(pattern)

	if noOfColumns > 0 {
		if fn, exists := cmdFunctions[pattern[0]]; exists {
			cook.CookConfig()
			fn(pattern[1:])
			os.Exit(0)
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
		columnCases = cook.UpdateCases(caseValue, noOfColumns)
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

	cook.VPrint(fmt.Sprintf("Pattern: %v \n", pattern))

	return params, pattern
}
