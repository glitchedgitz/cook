package methods

import (
	"strings"
)

func Upper(values []string, useless string, array *[]string) {
	for _, v := range values {
		*array = append(*array, strings.ToUpper(v))
	}
}

func Lower(values []string, useless string, array *[]string) {
	for _, v := range values {
		*array = append(*array, strings.ToLower(v))
	}
}

func Title(values []string, useless string, array *[]string) {
	for _, v := range values {
		*array = append(*array, strings.Title(v))
	}
}
