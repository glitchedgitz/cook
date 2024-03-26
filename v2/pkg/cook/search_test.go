package cook

import (
	"fmt"
	"strings"
	"testing"
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
	}

	for i, scenario := range scenarios {
		t.Run(fmt.Sprintf("s%d:%s", i, scenario.search), func(t *testing.T) {
			fmt.Println()
			dashes := strings.Repeat("-", 49-len(scenario.search))
			fmt.Printf("Search: %s %s\n", dashes, scenario.search)
			COOK := NewWithoutConfig()
			COOK.Search(scenario.search)
		})
	}

}
