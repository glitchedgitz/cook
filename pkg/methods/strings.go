package methods

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func Sort(values []string, useless string, array *[]string) {
	sort.Strings(values)
	*array = append(*array, values...)
}

func SortUnique(values []string, useless string, array *[]string) {

	tmp := make(map[string]bool)

	sort.Strings(values)

	for _, v := range values {
		if tmp[v] {
			continue
		}

		tmp[v] = true
		*array = append(*array, v)
	}
}

func Reverse(values []string, useless string, array *[]string) {
	for _, v := range values {
		runes := []rune(v)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		*array = append(*array, string(runes))
	}
}

func Replace(values []string, value string, array *[]string) {
	s := strings.SplitN(value, ":", 2)
	replaceText := s[0]
	replaceWith := s[1]
	for _, v := range values {
		*array = append(*array, strings.ReplaceAll(v, replaceText, replaceWith))
	}
}

func Split(values []string, split string, array *[]string) {
	for _, v := range values {
		*array = append(*array, strings.Split(v, split)...)
	}
}

func SplitIndex(values []string, value string, array *[]string) {

	vals := strings.SplitN(value, ":", 2)
	split := vals[0]
	index, err := strconv.Atoi(vals[1])

	if err != nil {
		log.Fatalln("Not Integer Value: "+values[1], err)
	}

	for _, v := range values {
		vv := strings.Split(v, split)
		if len(vv) >= index+1 {
			*array = append(*array, vv[index])
		}
	}
}
