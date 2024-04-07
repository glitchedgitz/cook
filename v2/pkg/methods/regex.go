package methods

import cook "github.com/glitchedgitz/cook/v2/pkg/config"

func  (m *Methods) Regex(values []string, regex string, array *[]string) {
	data := []byte{}
	for _, v := range values {
		data = append(data, []byte(v+"\n")...)
	}
	cook.FindRegex(data, regex, array)
}
