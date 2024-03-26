package cook

import (
	"fmt"
	"strings"
	"testing"
)

func compareTwoArrays(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestNew(t *testing.T) {
	// TODO: Write test cases for the New function
	scenarios := []struct {
		pattern        []string
		expectedResult []string
		expectedError  bool
	}{
		// Integer
		{[]string{"test"}, []string{"test"}, false},
		{[]string{"a,b,c"}, []string{"a", "b", "c"}, false},
		{[]string{"a,b,c"}, []string{"a", "b", "d"}, true},
		{[]string{"a,b,c", "1,2,3"}, []string{"a1", "a2", "a3", "b1", "b2", "b3", "c1", "c2", "c3"}, false},

		//Repeat Operators
		{[]string{"r**5"}, []string{"r", "r", "r", "r", "r"}, false},
		{[]string{"r*5"}, []string{"rrrrr"}, false},

		// Ranges
		{[]string{"1-5"}, []string{"1", "2", "3", "4", "5"}, false},
		{[]string{"a-z"}, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}, false},
		{[]string{"1-5", "a,b,c"}, []string{"1a", "1b", "1c", "2a", "2b", "2c", "3a", "3b", "3c", "4a", "4b", "4c", "5a", "5b", "5c"}, false},

		// Wordlist
		{[]string{"sec-amazon-api-gateway.txt"}, []string{"AmazonAPIGateway_5m3r4dmaP1", "AmazonAPIGateway_6qpee1cnq6"}, false},

		{[]string{"-p", "test,vest", "p"}, []string{"test", "vest"}, false},
	}

	for i, scenario := range scenarios {
		t.Run(fmt.Sprintf("s%d:%s", i, scenario.pattern), func(t *testing.T) {
			fmt.Printf("\nScenario: ------------------------------------------ %d\n", i)
			COOK := New(&COOK{
				Pattern: scenario.pattern,
			})
			COOK.Generate()
			fmt.Printf("Given Pattern:     %v\n", strings.Join(scenario.pattern, " "))
			fmt.Printf("Generated Pattern: %v\n", COOK.Final)
			fmt.Printf("Expected Pattern:  %v\n", scenario.expectedResult)
			fmt.Printf("Expected Error:    %v\n", scenario.expectedError)

			areSame := compareTwoArrays(COOK.Final, scenario.expectedResult)

			if !scenario.expectedError && !areSame {
				t.Fatalf("For Pattern %v Expected %v, got %v", scenario.pattern, scenario.expectedResult, COOK.Final)
			}
			if scenario.expectedError && areSame {
				t.Fatalf("For Pattern %v Expected %v, got %v", scenario.pattern, scenario.expectedResult, COOK.Final)
			}
		})
	}

}
