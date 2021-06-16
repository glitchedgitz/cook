package main

import (
	"fmt"
	"os"
	"strings"

	"cook/pkg/core"
	"cook/pkg/parse"
)

// var version = "1.6"

var total = 0

// var home, _ = os.UserHomeDir()
var otherCases = false
var columnCases = make(map[int]map[string]bool)

//Initializing with empty string, so loops will run for 1st column
var final = []string{""}

var (
	help             = parse.B("-h")
	verbose          = parse.B("-v")
	Min              = parse.I("-min")
	lineMode         = parse.B("-line-mode")
	showConfig       = parse.B("-config")
	configPath       = parse.S("-config-path")
	caseValue        = parse.S("-case")
	updateCacheFiles = parse.B("-update-all")
	l337             = parse.I("-1337")
)

var params = make(map[string]string)
var leetValues = make(map[string][]string)

func leetBegin() {
	leetValues["0"] = []string{"o", "O"}
	leetValues["1"] = []string{"i", "I", "l", "L"}
	leetValues["3"] = []string{"e", "E"}
	leetValues["4"] = []string{"a", "A"}
	leetValues["5"] = []string{"s", "S"}
	leetValues["6"] = []string{"b"}
	leetValues["7"] = []string{"t", "T"}
	leetValues["8"] = []string{"B"}
}

func parseInput() (map[string]string, []string) {

	parse.Help = core.Banner
	parse.Parse()

	if help {
		core.ShowHelp()
	}

	if showConfig {
		core.CookConfig()
		core.ShowConfig()
	}

	if updateCacheFiles {
		core.CookConfig()
		core.UpdateCache()
		os.Exit(0)
	}

	core.ConfigPath = configPath
	core.Verbose = verbose

	params = parse.UserDefinedFlags()
	pattern := parse.Args

	noOfColumns := len(pattern)

	if Min < 0 {

		Min = noOfColumns - 1
	} else {
		if Min > noOfColumns {
			fmt.Println("Err: Min is greator than no of columns")
			os.Exit(0)
		}
		Min -= 1
	}

	if caseValue != "" {
		columnCases = core.UpdateCases(caseValue, noOfColumns)
	}

	if l337 > -1 {
		if l337 > 1 {
			fmt.Println("Err: -1337 can be 0 or 1, 0 - Calm Mode & 1 - Angry Mode", l337)
			os.Exit(0)
		}
		leetBegin()
	}

	return params, pattern
}

func useless(s string) string {
	return s
}

func prefixSufixMode(values []string, array *[]string, fn func(string) string) {
	till := len(final)
	if len(final) > len(values) {
		till = len(values)
	}
	for i := 0; i < till; i++ {
		*array = append(*array, final[i]+fn(values[i]))
	}
}

func applyCase(values []string, array *[]string, fn func(string) string) {
	for _, t := range final {
		for _, v := range values {
			*array = append(*array, t+fn(v))
		}
	}
}

func applyColumnCases(columnValues []string, columnNum int, applyFunc func([]string, *[]string, func(string) string)) {
	temp := []string{}

	// Using cases for columnValues
	if len(columnCases[columnNum]) > 0 {
		otherCases = true
		//All cases
		if columnCases[columnNum]["A"] {

			applyFunc(columnValues, &temp, strings.ToUpper)
			applyFunc(columnValues, &temp, strings.ToLower)
			applyFunc(columnValues, &temp, strings.Title)

		} else {

			if !core.UpperCase && columnCases[columnNum]["U"] {
				applyFunc(columnValues, &temp, strings.ToUpper)
			}

			if columnCases[columnNum]["L"] {
				applyFunc(columnValues, &temp, strings.ToLower)
			}

			if columnCases[columnNum]["T"] {
				applyFunc(columnValues, &temp, strings.Title)
			}

		}

	} else {
		applyFunc(columnValues, &temp, useless)
	}

	final = temp
}

func printIt(fn func(string) string) {
	if l337 > -1 {
		for _, v := range final {
			v = fn(v)
			fmt.Println(v)
			v2 := v
			for l, ch := range leetValues {
				for _, c := range ch {
					if strings.Contains(v, c) {
						total++
						t := strings.ReplaceAll(v, c, l)
						v2 = strings.ReplaceAll(v2, c, l)
						if l337 == 1 {
							fmt.Println(t)
							if t != v2 {
								fmt.Println(v2)
							}
						}
					}
				}
			}
			if l337 == 0 {
				fmt.Println(v2)
			}
		}
	} else {
		// otherCases = true
		for _, v := range final {
			v = fn(v)
			total++
			fmt.Println(v)
		}
	}
}

func main() {
	// fmt.Fprintln(os.Stderr, banner)

	params, pattern := parseInput()

	core.CookConfig()

	for columnNum, param := range pattern {

		columnValues := []string{}
		var success = false

		for _, p := range strings.Split(param, ",") {

			success = core.ParseRanges(p, &columnValues)
			if success {
				continue
			}

			if val, exists := params[p]; exists {
				core.ParseValue(val, &columnValues)
				continue
			}

			// Raw String using `
			if strings.HasPrefix(p, "`") && strings.HasSuffix(p, "`") {
				lv := len(p)
				columnValues = append(columnValues, []string{p[1 : lv-1]}...)
				continue
			}

			success = core.CheckYaml(p, &columnValues)
			if success {
				continue
			}

			columnValues = append(columnValues, p)
		}

		if !lineMode || columnNum == 0 {
			applyColumnCases(columnValues, columnNum, applyCase)
		} else {
			applyColumnCases(columnValues, columnNum, prefixSufixMode)
		}

		if columnNum >= Min {
			if core.UpperCase {
				printIt(strings.ToUpper)
			}
			if core.LowerCase {
				printIt(strings.ToLower)
			}
			if (!core.LowerCase && !core.UpperCase) || otherCases {
				printIt(useless)
			}
		}
	}

	core.VPrint(fmt.Sprintf("Total words generated: %d", total))
}
