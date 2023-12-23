package main

import (
	"fmt"
	"strings"
)

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
