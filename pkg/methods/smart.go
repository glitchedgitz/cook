package methods

import "strings"

func SmartWords(words []string, useless string, array *[]string) {
	for _, word := range words {
		str := []string{}

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

		*array = append(*array, str...)
	}
}

func SmartWordsJoin(words []string, joinWith []string, array *[]string) {
	for _, word := range words {
		str := []string{}

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

		for _, join := range joinWith {
			*array = append(*array, strings.Join(str, join))
		}

	}
}
