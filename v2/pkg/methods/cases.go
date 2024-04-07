package methods

import (
	"strings"
)

func (m *Methods) Upper(values []string, useless string, array *[]string) {
	for _, v := range values {
		*array = append(*array, strings.ToUpper(v))
	}
}

func (m *Methods) Lower(values []string, useless string, array *[]string) {
	for _, v := range values {
		*array = append(*array, strings.ToLower(v))
	}
}

func (m *Methods) Title(values []string, useless string, array *[]string) {
	for _, v := range values {
		*array = append(*array, strings.Title(v))
	}
}
