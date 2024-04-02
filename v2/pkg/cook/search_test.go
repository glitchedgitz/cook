package cook

import (
	"fmt"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestSearch(t *testing.T) {
	// TODO: Write test cases for the New function
	scenarios := []struct {
		search string
	}{
		// Integer
		{"test"},
		{"api"},
		{"xss"},
		{"tld"},
		{"orange"},
	}

	for i, scenario := range scenarios {
		t.Run(fmt.Sprintf("s%d:%s", i, scenario.search), func(t *testing.T) {
			fmt.Println()
			dashes := strings.Repeat("-", 49-len(scenario.search))
			fmt.Printf("Search: %s %s\n", dashes, scenario.search)
			COOK := NewWithoutConfig()
			results, found := COOK.Search(scenario.search)
			if found {
				//Convert results object to yaml
				r, err := yaml.Marshal(results)
				if err != nil {
					t.Fatal("Error: ", err)
				}
				fmt.Println(string(r))
			} else {
				fmt.Println("Not Found")
			}
		})
	}

}
