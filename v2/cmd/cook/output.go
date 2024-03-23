package main

import (
	"fmt"
	"strings"

	cook "github.com/glitchedgitz/cook/v2/pkg/config"
)

func appendMode(values []string) {
	tmp := []string{}
	till := len(final)
	if len(final) > len(values) {
		till = len(values)
	}
	for i := 0; i < till; i++ {
		tmp = append(tmp, final[i]+values[i])
	}
	final = tmp
}

func permutationMode(values []string) {
	tmp := []string{}
	for _, t := range final {
		for _, v := range values {
			tmp = append(tmp, t+v)
		}
	}
	final = tmp
}

func checkParam(p string, array *[]string) bool {
	if val, exists := params[p]; exists {
		if cook.PipeInput(val, array) || cook.RawInput(val, array) || repeatOp(val, array) || cook.ParseFunc(val, array) || cook.ParseFile(p, val, array) || checkMethods(val, array) {
			return true
		}

		*array = append(*array, splitValues(val)...)
		return true
	}
	return false
}

func print() {

	if len(methodsForAll) > 0 {
		tmp := []string{}

		for _, meth := range strings.Split(methodsForAll, ",") {
			applyMethods(final, splitMethods(meth), &tmp)
		}
		for _, v := range tmp {
			fmt.Println(v)
		}
	} else {
		for _, v := range final {
			fmt.Println(v)
		}
	}
}
