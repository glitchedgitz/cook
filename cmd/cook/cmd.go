package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/giteshnxtlvl/cook/pkg/cook"
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
		fmt.Println("\nNot Found " + set + "\nTry charset, extensions, patterns, files, raw-files, ports or <file>.yaml")
	}
}
