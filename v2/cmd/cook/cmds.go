package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/glitchedgitz/cook/v2/pkg/config"
	"github.com/glitchedgitz/cook/v2/pkg/parse"
	"github.com/glitchedgitz/cook/v2/pkg/util"
)

func printWordlistNames(k, search string) string {
	return util.TerminalColor(strings.ReplaceAll(k,
		search, util.TerminalColor(search, util.Background)+util.Red), util.Red) + util.Reset
}

func searchMode(cmds []string) {
	found := false
	search := cmds[0]
	// convert the below fucntion to as data to map

	for cat, vv := range COOK.Config.M {
		for k, v := range vv {
			k = strings.ToLower(k)

			if strings.Contains(k, search) {
				if cat == "files" || cat == "raw-files" {

					coloredName := printWordlistNames(k, search)

					links := ""

					for i, file := range v {
						links += fmt.Sprintf(" " + util.TerminalLink(file, fmt.Sprintf("%d", i+1), util.Blue))
					}
					fmt.Printf("%-90s Links[%s ]", coloredName, links)
				} else if cat == "functions" {
					config.PrintFunc(k, v, search)
				} else {
					coloredName := printWordlistNames(k, search)
					// words := fmt.Sprintf(strings.ReplaceAll(fmt.Sprintf("    %v\n", v), search, util.Blue+search+config.Reset))
					words := util.TerminalColor(fmt.Sprint(v), util.Blue)
					fmt.Printf("%-90s Wordset %s", coloredName, words)
				}
				found = true
				fmt.Println()
			}
		}
	}

	if !found {
		fmt.Println("Not Found: ", search)
	}
	// COOK.SearchMode(cmds[0])
}

func showMode(cmds []string) {
	set := cmds[0]
	COOK.Show(set)
}

func addMode(cmds []string) {
	if len(cmds) != 3 {
		log.Println("Usage: cook add [keyword]=[values separated by comma] in [category]")
	}
	k := strings.SplitN(cmds[0], "=", 2)
	keyword := k[0]
	values := parse.SplitValues(k[1])
	category := cmds[2]
	COOK.Add(category, keyword, values)
}

func cleanMode(cmds []string) {
	fmt.Println("Not implemented yet")
	COOK.Clean()
}

func infoMode(cmds []string) {
	set := cmds[0]
	COOK.Info(set)
}

func updateMode(cmds []string) {
	f := cmds[0]
	COOK.Update(f)

}

func deleteMode(cmds []string) {
	if len(cmds) != 1 {
		log.Fatalln("Usage: cook delete [keyword]")
	}
	keyword := cmds[0]
	COOK.Delete(keyword)
}
