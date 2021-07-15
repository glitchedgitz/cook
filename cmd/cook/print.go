package main

import (
	"fmt"

	"github.com/giteshnxtlvl/cook/pkg/cook"
)

func useless(s string) string {
	return s
}

var doLeet = false
var doEncode = false

func print() {
	tmp := []string{}
	finalClone := final
	if cook.UpperCase {
		cook.Cases(finalClone, []string{"U"}, &tmp)
		finalClone = tmp
		tmp = nil
	}
	if cook.LowerCase {
		cook.Cases(finalClone, []string{"L"}, &tmp)
		finalClone = tmp
		tmp = nil
	}
	// if cook.LowerCase {
	// 	printIt(strings.ToLower)
	// }
	// if (!cook.LowerCase && !cook.UpperCase) || otherCases {
	// 	printIt(useless)
	// }
	if doLeet {
		cook.Leet(finalClone, l337, &tmp)
		finalClone = tmp
		tmp = nil
	}
	if doEncode {
		cook.Encode(finalClone, encodeString, &tmp)
		finalClone = tmp
		tmp = nil
	}
	for _, v := range finalClone {
		fmt.Println(v)
	}
}
