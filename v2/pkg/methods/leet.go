package methods

import (
	"log"
	"strconv"
	"strings"
)

func (m *Methods) Leet(values []string, value string, array *[]string) {

	mode, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalln("Err: Leet can be 0 or 1, usage: leet[0]")
	}

	for _, v := range values {
		var tmp = make(map[string]bool)
		v2 := v
		for l, ch := range m.LeetValues {
			for _, c := range ch {
				if strings.Contains(v, c) {
					t := strings.ReplaceAll(v, c, l)
					v2 = strings.ReplaceAll(v2, c, l)
					tmp[t] = true
					tmp[v2] = true
				}
			}
		}

		if mode == 0 {
			*array = append(*array, v2)
		} else {
			for k := range tmp {
				*array = append(*array, k)
			}
		}
	}
}
