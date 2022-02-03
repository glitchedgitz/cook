package main

import (
	"fmt"
)

func print() {

	if len(methodsForAll) > 0 {
		tmp := []string{}
		applyMethods(final, splitMethods(methodsForAll), &tmp)
		for _, v := range tmp {
			fmt.Println(v)
		}
	} else {
		for _, v := range final {
			fmt.Println(v)
		}
	}
}
