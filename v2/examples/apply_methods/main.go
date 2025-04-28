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

	// Print the original strings
	fmt.Println("Original strings:")
	for _, s := range inputStrings {
		fmt.Println("  -", s)
	}

	// Try with a non-existent method
	fmt.Println("\nTrying with non-existent method:")
	methods := []string{"upper", "reverse", "uppercase"}
	result, err := generator.ApplyMethods(inputStrings, methods)
	if err != nil {
		// For method-not-found errors, print the error and continue with original strings
		fmt.Println("Note:", err)
		result = inputStrings
	}

	// Print the transformed strings
	fmt.Println("\nAfter applying methods (upper, reverse, uppercase):")
	for _, s := range result {
		fmt.Println("  -", s)
	}

	// Try with base64 decoding on invalid input
	fmt.Println("\nTrying base64 decode on invalid input:")
	invalidInput := []string{"not-a-base64-string"}
	decoded, err := generator.ApplyMethods(invalidInput, []string{"b64d"})
	if err != nil {
		// For method-not-found errors, print the error and continue with original strings
		fmt.Println("Note:", err)
		decoded = invalidInput
	}
	for _, s := range decoded {
		fmt.Println("  -", s)
	}
}
