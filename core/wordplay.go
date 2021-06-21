package core

import "strings"

func WordPlay(words []string, joinWith string, fn func(string) string, array *[]string) {

	for _, word := range words {

		str := []string{}
		w := ""

		if strings.Contains(word, "_") {
			str = strings.Split(word, "_")

		} else if strings.Contains(word, "-") {
			str = strings.Split(word, "-")

		} else {

			j := 0
			for i, letter := range word {
				if letter > 'A' && letter < 'Z' {
					str = append(str, word[j:i])
					j = i
				}
			}
			str = append(str, word[j:])
		}

		last := len(str) - 1
		if len(str) > 1 {
			for _, s := range str[:last] {
				w += fn(s) + joinWith
			}
		}
		w += fn(str[last])

		*array = append(*array, w)

	}

}
