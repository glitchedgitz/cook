package cook

import (
	"strings"
)

func (cook *COOK) Search(search string) (map[string]map[string][]string, bool) {

	searchMap := make(map[string]map[string][]string)

	found := true

	for cat, vv := range cook.Config.Ingredients {
		for k, v := range vv {
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
