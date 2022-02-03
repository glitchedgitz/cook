package methods

import "github.com/giteshnxtlvl/cook/v2/pkg/cook"

func Regex(values []string, regex string, array *[]string) {
	data := []byte{}
	for _, v := range values {
		data = append(data, []byte(v+"\n")...)
	}
	cook.FindRegex(data, regex, array)
}
