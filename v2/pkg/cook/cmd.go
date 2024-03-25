package cook

import (
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/glitchedgitz/cook/v2/pkg/config"
	"github.com/glitchedgitz/cook/v2/pkg/util"
	"github.com/manifoldco/promptui"
)

const myyaml = "my.yaml"

func (cook *COOK) analyseParams(params map[string]string) {
	for param, value := range params {
		// fmt.Println(params)
		if strings.HasSuffix(param, ":") {
			delete(params, param)
			param = strings.TrimSuffix(param, ":")
			cook.Config.InputFile[param] = true
			params[param] = value
		}
	}
}

// Add new set in custom.yaml
// cook add [keyord]=[values separated by comma] in [category]
func (cook *COOK) Add(category, keyword string, values []string) {

	m := make(map[string]map[string][]string)
	util.ReadYaml(path.Join(cook.Config.IngredientsPath, myyaml), m)

	if _, exists := m[category]; exists {
		m[category][keyword] = append(m[category][keyword], values...)
	} else {
		m[category] = map[string][]string{
			keyword: values,
		}
	}

	util.WriteYaml(path.Join(cook.Config.IngredientsPath, myyaml), m)
	fmt.Printf("Added \"%s\" in \"%s\" ", keyword, category)
}

func (cook *COOK) Update(f string) {

	const (
		updateAllStr   = "*"
		updateDBStr    = "db"
		updateCacheStr = "cache"
	)

	if f == updateAllStr {
		cook.Config.UpdateDb()
		cook.Config.UpdateCache()
	} else if f == updateDBStr {
		cook.Config.UpdateDb()
	} else if f == updateCacheStr {
		cook.Config.UpdateCache()
	} else if files, exists := cook.Config.M["files"][f]; exists {
		filepath := path.Join(cook.Config.CachePath, f)
		os.Remove(filepath)
		cook.Config.CheckFileCache(f, files)
	} else {
		log.Println("No mode or keyword found\nUse \"db\" to update cook-ingredients\nUse \"cache\" to update cached files from soruce\nUse \"*\" to do the both")
	}
}

// Delete from custom.yaml
// cook delete [keyword]
func (cook *COOK) Delete(keyword string) {

	m := make(map[string]map[string][]string)
	util.ReadYaml(myyaml, m)
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
		util.WriteYaml(path.Join(cook.Config.IngredientsPath, myyaml), m)
	} else {
		log.Fatalln("Keyword doesn't exists")
	}

}

func (cook *COOK) Clean() {
}

func (cook *COOK) Info(set string) {

	m := make(map[string]map[string][]string)
	if strings.HasSuffix(set, ".yaml") || strings.HasPrefix(set, ".yml") {
		util.ReadYaml(set, m)
	}

	fmt.Println("\n" + util.Blue + set + util.Reset)
	fmt.Println("Path    : ", path.Join(cook.Config.IngredientsPath, set))
	fmt.Println("Sets    : ", len(m))
	fmt.Println("Version : ", len(m))
}

func (cook *COOK) Show(set string) {

	if strings.HasSuffix(set, ".yaml") || strings.HasPrefix(set, ".yml") {
		fmt.Println(string(util.ReadFile(set)))
		return
	}

	if vals, exists := cook.Config.M[set]; exists {
		fmt.Printf("\n" + util.Blue + strings.ToUpper(set) + util.Reset + "\n\n")

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
				config.PrintFunc(k, vals[k], "")
			}
		}
	} else {
		fmt.Println("\nNot Found " + set + "\nTry functions, files, raw-files, ports or <file>.yaml")
	}
}
