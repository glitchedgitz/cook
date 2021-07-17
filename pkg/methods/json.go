package methods

import (
	"strings"

	"github.com/buger/jsonparser"
)

func GetJsonField(lines []string, get string, array *[]string) {
	values := strings.Split(get, ":")
	for _, line := range lines {
		data := []byte(line)
		value, _, _, _ := jsonparser.Get(data, values...)
		v := string(value)
		*array = append(*array, v)
	}
}
