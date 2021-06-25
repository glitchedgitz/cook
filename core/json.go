package core

import (
	"github.com/buger/jsonparser"
)

func GetJsonField(lines []string, get []string, array *[]string) {
	for _, line := range lines {
		data := []byte(line)
		value, _, _, err := jsonparser.Get(data, get...)
		if err != nil {
			panic(err)
		}
		// fmt.Println(string(value))
		v := string(value)
		*array = append(*array, v)
	}
}
