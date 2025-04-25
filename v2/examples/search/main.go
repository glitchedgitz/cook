package main

import (
	"fmt"

	"github.com/glitchedgitz/cook/v2/pkg/cook"
)

func main() {
	// Initialize the generator once with all the shared resources
	generator := cook.NewGenerator()

	// Use the generator for search operations
	searchTerm := "api"
	fmt.Printf("Searching for: '%s'\n", searchTerm)

	// Search through the Cook database using the generator
	results, found := generator.Search(searchTerm)

	if !found {
		fmt.Printf("No results found for '%s'\n", searchTerm)
		return
	}

	fmt.Printf("Found %d results:\n\n", len(results))

	// Display the search results
	for i, result := range results {
		fmt.Printf("%d. %s (%s)\n", i+1, result.Name, result.Type)

		if result.Path != "" {
			fmt.Printf("   Path: %s\n", result.Path)
		}

		// Display a sample of the content
		if len(result.Content) > 0 {
			sampleSize := min(3, len(result.Content))
			fmt.Printf("   Content sample: %v\n", result.Content[:sampleSize])
			if len(result.Content) > sampleSize {
				fmt.Printf("   ... and %d more items\n", len(result.Content)-sampleSize)
			}
		}

		fmt.Println()
	}

	// Example of using generator for both search and generation
	fmt.Println("=== Using the same generator for pattern generation ===")

	// Generate patterns with the same generator
	patterns := []string{"admin", "2023"}
	wordlist := generator.Generate(patterns)

	fmt.Printf("Generated %d patterns from %v\n", len(wordlist), patterns)
	if len(wordlist) > 0 {
		sampleSize := min(5, len(wordlist))
		fmt.Printf("Sample: %v\n", wordlist[:sampleSize])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
