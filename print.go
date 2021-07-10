package main

import (
	"fmt"

	"github.com/giteshnxtlvl/cook/core"
)

func useless(s string) string {
	return s
}

var doLeet = false
var doEncode = false

func print() {
	tmp := []string{}
	finalClone := final
	if core.UpperCase {
		core.Cases(finalClone, []string{"U"}, &tmp)
		finalClone = tmp
		tmp = nil
	}
	if core.LowerCase {
		core.Cases(finalClone, []string{"L"}, &tmp)
		finalClone = tmp
		tmp = nil
	}
	// if core.LowerCase {
	// 	printIt(strings.ToLower)
	// }
	// if (!core.LowerCase && !core.UpperCase) || otherCases {
	// 	printIt(useless)
	// }
	if doLeet {
		core.Leet(finalClone, l337, &tmp)
		finalClone = tmp
		tmp = nil
	}
	if doEncode {
		core.Encode(finalClone, encodeString, &tmp)
		finalClone = tmp
		tmp = nil
	}
	for _, v := range finalClone {
		fmt.Println(v)
	}
}
