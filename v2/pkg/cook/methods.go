package cook

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/glitchedgitz/cook/v2/pkg/parse"
)

// ValidateMethod checks if a method exists and returns an error if it doesn't
func (cook *COOK) ValidateMethod(method string) error {
	// Extract method name and parameters
	methodName := method
	if strings.Contains(method, "[") {
		methodName, _ = parse.ReadSqBr(method)
	}

	// Check if method exists in any of the method maps
	_, methodExists := cook.Method.MethodFuncs[methodName]
	_, urlExists := cook.Method.UrlFuncs[methodName]
	_, encoderExists := cook.Method.EncodersFuncs[methodName]

	if !methodExists && !urlExists && !encoderExists {
		suggestions := cook.MistypedCheck(methodName)
		return fmt.Errorf("method '%s' not found. \n%s", methodName, suggestions)
	}

	return nil
}

// ApplyMethods applies the given methods to the input strings and handles errors gracefully
func (cook *COOK) ApplyMethods(input []string, methods []string, output *[]string) error {
	// Validate all methods first
	for _, method := range methods {
		if err := cook.ValidateMethod(method); err != nil {
			return err
		}
	}

	// If all methods are valid, apply them
	tmp := input
	for _, method := range methods {
		methodName := method
		methodParam := ""
		if strings.Contains(method, "[") {
			methodName, methodParam = parse.ReadSqBr(method)
		}

		// Try each type of method
		if fn, exists := cook.Method.MethodFuncs[methodName]; exists {
			fn(tmp, methodParam, output)
			tmp = *output
			*output = []string{}
			continue
		}

		if fn, exists := cook.Method.UrlFuncs[methodName]; exists {
			for _, t := range tmp {
				u, err := url.Parse(t)
				if err != nil {
					continue
				}
				fn(u, output)
			}
			tmp = *output
			*output = []string{}
			continue
		}

		if encoder, exists := cook.Method.EncodersFuncs[methodName]; exists {
			for _, t := range tmp {
				encoded, err := encoder.Encode([]byte(t))
				if err != nil {
					*output = append(*output, fmt.Sprintf("error in %s: %v", methodName, err))
					continue
				}
				*output = append(*output, string(encoded))
			}
			tmp = *output
			*output = []string{}
		}
	}

	*output = tmp
	return nil
}

func (cook *COOK) MistypedCheck(mistyped string) string {
	var suggestions []string
	check := func(k string) {
		similarity := strutil.Similarity(mistyped, k, metrics.NewHamming())
		if similarity >= 0.3 {
			suggestions = append(suggestions, k)
		}
	}

	for k := range cook.Method.MethodFuncs {
		check(k)
	}

	for k := range cook.Method.UrlFuncs {
		check(k)
	}

	for k := range cook.Method.EncodersFuncs {
		check(k)
	}

	if len(suggestions) == 0 {
		return "No similar methods found"
	}

	return "Similar methods: " + strings.Join(suggestions, "\n - ")
}

// CheckMethods checks if a string contains methods and applies them
func (cook *COOK) CheckMethods(value string, array *[]string) bool {
	if strings.Contains(value, ".") {
		methods := parse.SplitMethods(value)
		tmp := []string{}
		if err := cook.ApplyMethods([]string{methods[0]}, methods[1:], &tmp); err != nil {
			fmt.Printf("Error: %v\n", err)
			return false
		}
		*array = append(*array, tmp...)
		return true
	}
	return false
}
