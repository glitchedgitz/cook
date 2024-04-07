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

	for cat, vv := range COOK.Config.Ingredients {
		for k, v := range vv {
			k = strings.ToLower(k)
			k = strings.TrimSpace(k)
			if cat == "files" || cat == "raw-files" {
				link := ""
				if strings.HasPrefix(v[0], "https://raw.githubusercontent.com") {
					path := strings.Split(v[0], "/")[4:]
					link = strings.ToTitle(path[0]) + " > " + strings.Join(path[2:len(path)-1], " > ")
				} else {
					d := strings.TrimPrefix(v[0], "http://")
					d = strings.TrimPrefix(d, "https://")
					link = d
					// link = strings.Join(strings.Split(d, "/"), " > ")
				}
				if strings.Contains(strings.ToLower(link+k), search) {
					coloredName := k
					coloredName = printWordlistNames(k, search)
					link = strings.ToLower(link)

					link = printWordlistNames(link, search)
					links := ""
					for i, file := range v {

						// fmt.Println(link)
						links += fmt.Sprintf(" " + util.TerminalLink(file, fmt.Sprintf("%d", i+1), util.Blue))
					}

					// because of color encoding using %-70s was not working right
					repeat := 50 - len(k)
					extraSpace := ""
					if repeat > 0 {
						extraSpace = strings.Repeat(" ", repeat)
					}

					fmt.Printf("%s%s  %s [%s ]", coloredName, extraSpace, link, links)
					found = true
					fmt.Println()
				}
			} else {
				if strings.Contains(k, search) {
					if cat == "functions" {
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
