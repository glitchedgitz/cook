package main

import (
	"fmt"
	"strings"

	"github.com/glitchedgitz/cook/v2/pkg/cook"
)

func printWordlistNames(k, search string) string {
	return cook.TerminalColor(strings.ReplaceAll(k,
		search, cook.TerminalColor(search, cook.Background)+cook.Red), cook.Red) + cook.Reset
}

func searchMode(cmds []string) {

	search := cmds[0]
	found := false

	for cat, vv := range cook.M {
		for k, v := range vv {
			k = strings.ToLower(k)

			if strings.Contains(k, search) {
				if cat == "files" || cat == "raw-files" {

					coloredName := printWordlistNames(k, search)

					links := ""

					for i, file := range v {
						links += fmt.Sprintf(" " + cook.TerminalLink(file, fmt.Sprintf("%d", i+1), cook.Blue))
					}
					fmt.Printf("%-90s Links[%s ]", coloredName, links)
				} else if cat == "functions" {
					cook.PrintFunc(k, v, search)
				} else {
					coloredName := printWordlistNames(k, search)
					// words := fmt.Sprintf(strings.ReplaceAll(fmt.Sprintf("    %v\n", v), search, cook.Blue+search+cook.Reset))
					words := cook.TerminalColor(fmt.Sprint(v), cook.Blue)
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
}
