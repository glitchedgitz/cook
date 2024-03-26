package cook

import (
	"fmt"
	"strings"
)

func (cook *COOK) Search(search string) (map[string]map[string][]string, bool) {

	searchMap := make(map[string]map[string][]string)

	found := true

	for cat, vv := range cook.Config.M {
		for k, v := range vv {
			fmt.Println(cat, k, v)
			k = strings.ToLower(k)
			if strings.Contains(k, search) {
				if !found {
					found = true
				}
				searchMap[cat] = map[string][]string{k: v}
			}
		}
	}

	return searchMap, found

}
