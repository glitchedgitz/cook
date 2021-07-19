package methods

import "strconv"

// Get Charcode of string
func Charcode(values []string, semicolon string, array *[]string) {
	addSemiColon := false
	if semicolon == "1" {
		addSemiColon = true
	}

	for _, str := range values {
		encoded := ""
		for _, s := range str {
			encoded += "&#" + strconv.Itoa(int(s))
			if addSemiColon {
				encoded += ";"
			}
		}
		*array = append(*array, encoded)
	}
}
