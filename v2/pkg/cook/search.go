package cook

import (
	"strings"
)

type CookIngredient struct {
	// keyword name
	Name string `json:"name"`

	// Type: files, raw-files, functions, sets, etc
	Type string `json:"type"`

	// Path: only for files and raw-files, otherwise empty
	Path string `json:"path"`

	// Content
	// - For files and raw-files it will be `links`
	// - For functions it will be permutatons
	// - For sets it will be the wordsets
	// - For ports it will be the ports
	Content []string `json:"content"`
}

func (cook *COOK) Search(search string) ([]CookIngredient, bool) {
	found := false
	// convert the below fucntion to as data to map
	var searches = []CookIngredient{}

	for category, item := range cook.Config.Ingredients {
		for keyword, content := range item {

			keyword = strings.ToLower(keyword)
			keyword = strings.TrimSpace(keyword)

			if category == "files" || category == "raw-files" {
				link := ""
				if strings.HasPrefix(content[0], "https://raw.githubusercontent.com") {
					path := strings.Split(content[0], "/")[4:]
					link = strings.ToTitle(path[0]) + " > " + strings.Join(path[2:len(path)-1], " > ")
				} else {
					d := strings.TrimPrefix(content[0], "http://")
					d = strings.TrimPrefix(d, "https://")
					link = d
					// link = strings.Join(strings.Split(d, "/"), " > ")
				}

				if strings.Contains(strings.ToLower(link+keyword), search) {
					link = strings.ToLower(link)
					searches = append(searches,
						CookIngredient{
							Name:    keyword,
							Type:    category,
							Path:    link,
							Content: content,
						},
					)
					found = true
				}
			} else {
				if strings.Contains(keyword, search) {
					searches = append(searches,
						CookIngredient{
							Name:    keyword,
							Type:    category,
							Path:    "",
							Content: content,
						},
					)
					found = true
				}
			}
		}
	}

	return searches, found
}
