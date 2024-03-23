package main

import (
	"fmt"
	"time"

	cook "github.com/glitchedgitz/cook/v2/pkg/config"
)

func run() {
	for columnNum, param := range pattern {

		columnValues := []string{}

		for _, p := range splitValues(param) {
			cook.VPrint(fmt.Sprintf("Param: %s \n", p))
			if cook.RawInput(p, &columnValues) || cook.ParseRanges(p, &columnValues) || repeatOp(p, &columnValues) || checkMethods(p, &columnValues) || checkParam(p, &columnValues) || cook.CheckYaml(p, &columnValues) {
				continue
			}
			columnValues = append(columnValues, p)
		}

		cook.VPrint(fmt.Sprintf("%-40s: %s", "Time after getting values", time.Since(start)))

		if mapval, exists := methodMap[columnNum]; exists {
			tmp := []string{}
			applyMethods(columnValues, mapval, &tmp)
			columnValues = tmp
		}

		if !appendMap[columnNum] || columnNum == 0 {
			permutationMode(columnValues)
		} else {
			appendMode(columnValues)
		}

		cook.VPrint(fmt.Sprintf("%-40s: %s", "Time ApplyColumnCases", time.Since(start)))

		if columnNum >= min {
			print()
		}
	}
}
