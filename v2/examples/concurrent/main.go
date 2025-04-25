package main

import (
	"fmt"
	"sync"

	"github.com/glitchedgitz/cook/v2/pkg/cook"
)

func main() {
	// Initialize the generator once with all the shared resources
	generator := cook.NewGenerator()

	// Define multiple patterns to generate with various methods
	patterns := [][]string{
		{"a,b,c", "1,2,3"}, // Basic pattern
		{"1-10"},           // Range pattern
		{"admin,Admin", "2023-2025", "-m", "leet[0]"},   // Multiple values
		{"admin", "password", "-m", "upper"},            // All uppercase
		{"admin", "password", "-m", "md5"},              // MD5 hash
		{"test", "test*5", "-m", "b64e"},                // Base64 encoding
		{"adminUser", "admin-api", "-m", "smart"},       // Smart breakdown
		{"admin", "secret", "-mc", "0:upper;1:reverse"}, // Column-specific methods
	}

	// Channel to collect results
	resultsChan := make(chan []string, len(patterns))

	// WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup
	wg.Add(len(patterns))

	// Generate patterns concurrently
	for i, pattern := range patterns {
		// Create a separate copy of the pattern for each goroutine
		patternCopy := make([]string, len(pattern))
		copy(patternCopy, pattern)

		go func(index int, p []string) {
			defer wg.Done()

			results := generator.Generate(p)
			resultsChan <- results

			// Describe what was generated
			description := p
			if len(description) > 3 {
				description = description[:2] // Just show first 2 elements if pattern has method flags
				description = append(description, "...")
			}

			fmt.Printf("Pattern #%d %v generated %d results\n", index+1, description, len(results))
			if len(results) > 0 {
				// Show a sample of the results
				sampleSize := min(5, len(results))
				fmt.Printf("  Sample: %v\n\n", results[:sampleSize])
			}
		}(i, patternCopy)
	}

	// Close the channel when all goroutines are done
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Collect all results
	allResults := []string{}
	for results := range resultsChan {
		allResults = append(allResults, results...)
	}

	fmt.Printf("\nTotal combined results: %d\n", len(allResults))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
