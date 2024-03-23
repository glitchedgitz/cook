package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	cook "github.com/glitchedgitz/cook/v2/pkg/config"
	"github.com/manifoldco/promptui"
)

var myyaml = "my.yaml"

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
	cook.ReadYaml(myyaml, m)

	if _, exists := m[category]; exists {
		m[category][keyword] = append(m[category][keyword], values...)
	} else {
		m[category] = map[string][]string{
			keyword: values,
		}
	}

	cook.WriteYaml(path.Join(cook.ConfigFolder, cook.IngredientsFolder, myyaml), m)
	fmt.Printf("Added \"%s\" in \"%s\" ", keyword, category)
}

func updateMode(cmds []string) {
	f := cmds[0]

	if f == "*" {
		cook.UpdateDb()
		cook.UpdateCache()
	} else if f == "db" {
		cook.UpdateDb()
	} else if f == "cache" {
		cook.UpdateCache()
	} else if files, exists := cook.M["files"][f]; exists {
		filepath := path.Join(home, ".cache", "cook", f)
		os.Remove(filepath)
		cook.CheckFileCache(f, files)
	} else {
		log.Println("No mode or keyword found\nUse \"db\" to update cook-ingredients\nUse \"cache\" to update cached files from soruce\nUse \"*\" to do the both")
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
	cook.ReadYaml(myyaml, m)
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
		cook.WriteYaml(path.Join(cook.ConfigFolder, cook.IngredientsFolder, myyaml), m)
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
	fmt.Println("Path    : ", path.Join(cook.ConfigFolder, cook.IngredientsFolder, set))
	fmt.Println("Sets    : ", len(m))
	fmt.Println("Version : ", len(m))
}

func showMode(cmds []string) {
	set := cmds[0]

	if strings.HasSuffix(set, ".yaml") || strings.HasPrefix(set, ".yml") {
		fmt.Println(string(cook.ReadFile(path.Join(cook.ConfigFolder, cook.IngredientsFolder, set))))
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
		fmt.Println("\nNot Found " + set + "\nTry functions, files, raw-files, ports or <file>.yaml")
	}
}
