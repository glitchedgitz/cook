package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/glitchedgitz/cook/v2/pkg/parse"
)

func searchMode(cmds []string) {
	COOK.SearchMode(cmds[0])
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
