package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/giteshnxtlvl/cook/pkg/cook"
	"github.com/manifoldco/promptui"
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
				} else if cat == "functions" {
					cook.PrintFunc(k, v, search)
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

// Add new set in custom.yaml
// cook add [keyord]=[values separated by comma] in [category]
func addMode(cmds []string) {
	if len(cmds) != 3 {
		log.Println("Usage: cook add [keyword]=[values separated by comma] in [category]")
	}
	k := strings.SplitN(cmds[0], "=", 2)
	keyword := k[0]
	values := splitValues(k[1])
	category := cmds[2]
	m := make(map[string]map[string][]string)
	cook.ReadYaml("custom.yaml", m)

	if _, exists := m[category]; exists {
		m[category][keyword] = append(m[category][keyword], values...)
	} else {
		m[category][keyword] = values
	}

	cook.WriteYaml("custom.yaml", m)
	fmt.Printf("Added \"%s\" in \"%s\" ", keyword, category)
}

func updateMode(cmds []string) {
	filename := cmds[0]
	filepath := path.Join(home, ".cache", "cook", filename)
	if files, exists := cook.M["files"][filename]; exists {
		os.Remove(filepath)
		cook.CheckFileCache(filename, files)
	}
}

// Delete from custom.yaml
// cook delete [keyword]
func deleteMode(cmds []string) {
	if len(cmds) != 1 {
		log.Fatalln("Usage: cook delete [keyword]")
	}
	keyword := cmds[0]

	m := make(map[string]map[string][]string)
	cook.ReadYaml("custom.yaml", m)
	category := ""

	found := false
	for k, v := range m {

		if _, exists := v[keyword]; exists {
			category = k
			prompt := promptui.Select{
				Label: fmt.Sprintf("Are you sure, you want to delete \"%s\" from \"%s\"?", keyword, k),
				Items: []string{"No", "Yes"},
			}

			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			if result == "Yes" {
				found = true
			} else {
				log.Fatalln("Not deleted")
			}
			break
		}
	}

	if found {
		delete(m[category], keyword)
		fmt.Printf("Deleted \"%s\" from \"%s\" ", keyword, category)
		cook.WriteYaml("custom.yaml", m)
	} else {
		log.Fatalln("Keyword doesn't exists")
	}

}
func cleanMode(cmds []string) {
}
func infoMode(cmds []string) {
	set := cmds[0]

	m := make(map[string]map[string][]string)
	if strings.HasSuffix(set, ".yaml") || strings.HasPrefix(set, ".yml") {
		cook.ReadYaml(set, m)
	}

	fmt.Println("\n" + cook.Blue + set + cook.Reset)
	fmt.Println("Path    : ", path.Join(cook.ConfigFolder, "yaml", set))
	fmt.Println("Sets    : ", len(m))
	fmt.Println("Version : ", len(m))
}

func showMode(cmds []string) {
	set := cmds[0]

	if strings.HasSuffix(set, ".yaml") || strings.HasPrefix(set, ".yml") {
		fmt.Println(string(cook.ReadFile(path.Join(cook.ConfigFolder, "yaml", set))))
		return
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
				cook.PrintFunc(k, vals[k], "")
			}
		}
	} else {
		fmt.Println("\nNot Found " + set + "\nTry charset, extensions, patterns, files, raw-files, ports or <file>.yaml")
	}
}
