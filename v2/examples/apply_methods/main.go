package main

import (
	"fmt"

	"github.com/glitchedgitz/cook/v2/pkg/cook"
)

func main() {
	// Create a new generator
	generator := cook.NewGenerator()

	// Sample input strings
	inputStrings := []string{"test", "hello", "world"}

	// Define methods to apply
	// Available methods will depend on what's configured in your cook installation
	methods := []string{"upper", "reverse"}

	// Apply the methods to the input strings
	result := generator.ApplyMethods(inputStrings, methods)

	// Print the original strings
	fmt.Println("Original strings:")
	for _, s := range inputStrings {
		fmt.Println("  -", s)
	}

	// Print the transformed strings
	fmt.Println("\nAfter applying methods (upper, reverse):")
	for _, s := range result {
		fmt.Println("  -", s)
	}

	// Example with other methods
	fmt.Println("\nApplying different methods (leet, lower):")
	otherResult := generator.ApplyMethods(inputStrings, []string{"leet[0]", "lower"})
	for _, s := range otherResult {
		fmt.Println("  -", s)
	}
}
