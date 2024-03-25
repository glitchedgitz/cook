package cook

import (
	"fmt"
	"strings"

	"github.com/glitchedgitz/cook/v2/pkg/config"
	"github.com/glitchedgitz/cook/v2/pkg/util"
)

func printWordlistNames(k, search string) string {
	return util.TerminalColor(strings.ReplaceAll(k,
		search, util.TerminalColor(search, util.Background)+util.Red), util.Red) + util.Reset
}

func (cook *COOK) SearchMode(search string) {

	found := false

	for cat, vv := range cook.Config.M {
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
}
