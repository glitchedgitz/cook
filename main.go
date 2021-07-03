package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/giteshnxtlvl/cook/core"
	"github.com/giteshnxtlvl/cook/parse"

	"github.com/giteshnxtlvl/pencode/pkg/pencode"
)

var total = 0

// var home, _ = os.UserHomeDir()
var otherCases = false
var columnCases = make(map[int]map[string]bool)

//Initializing with empty string, so loops will run for 1st column
var final = []string{""}

var (
	help          = parse.B("-h", "-help")
	verbose       = parse.B("-v", "-verbose")
	showCol       = parse.B("-c", "-col")
	Min           = parse.I("-m", "-min")
	appendColumns = parse.S("-a", "-append")
	showConfig    = parse.B("-conf", "-config")
	caseValue     = parse.S("-ca", "-case")
	encodeValue   = parse.S("-e", "-encode")
	update        = parse.S("-u", "-update")
	l337          = parse.I("-l", "-leet")
)

var params = make(map[string]string)
var leetValues = make(map[string][]string)
var appendMode = make(map[int]bool)
var encodeString []string

var finalFunc = func(s string) {
	fmt.Println(s)
}

var justAppend = func(array *[]string, s string) {
	*array = append(*array, s)
}

var justPrint = func(array *[]string, s string) {
	fmt.Println(s)
}

var justEnd = justAppend

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

func encode(inputdata string) {
	chain := pencode.NewChain()
	err := chain.Initialize(encodeString)
	if err != nil {
		log.Fatalln(err)
	}
	output, err := chain.Encode([]byte(inputdata))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(output))
}

func printIt(fn func(string) string) {
	if l337 > -1 {
		for _, v := range final {
			v = fn(v)
			finalFunc(v)
			v2 := v
			for l, ch := range leetValues {
				for _, c := range ch {
					if strings.Contains(v, c) {
						total++
						t := strings.ReplaceAll(v, c, l)
						v2 = strings.ReplaceAll(v2, c, l)
						if l337 == 1 {
							finalFunc(t)
							if t != v2 {
								finalFunc(v2)
							}
						}
					}
				}
			}
			if l337 == 0 {
				finalFunc(v2)
			}
		}
	} else {
		// otherCases = true
		for _, v := range final {
			v = fn(v)
			total++
			finalFunc(v)
		}
	}
}

func checkParamAndYaml(p string, array *[]string) {
	if val, exists := params[p]; exists {
		core.ParseValue(p, val, array)
		return
	}

	if core.CheckYaml(p, array) {
		return
	}

	*array = append(*array, p)
}

func init() {
	log.SetFlags(0)
}

func main() {
	// start := time.Now()
	// fmt.Fprintln(os.Stderr, banner)
	params, pattern := parseInput()

	core.CookConfig()

	for columnNum, param := range pattern {

		columnValues := []string{}

		for _, p := range strings.Split(param, ",") {
			core.VPrint(fmt.Sprintf("Param: %s \n", p))

			if core.ParseRanges(p, &columnValues) {
				continue
			}

			// Raw String using `
			if strings.HasPrefix(p, "`") && strings.HasSuffix(p, "`") {
				lv := len(p)
				columnValues = append(columnValues, []string{p[1 : lv-1]}...)
				continue
			}

			// Checking for url func/encoding/json
			if strings.Count(p, ".") > 0 {
				splitS := strings.Split(p, ".")
				u := splitS[0]
				if _, exists := params[u]; exists {

					get := splitS[1:]

					tmp := []string{}
					vallll := []string{}

					checkParamAndYaml(u, &vallll)

					for _, g := range get {

						if g == "wordplay" {
							core.WordPlay(vallll, "*", useless, &tmp)
						} else if g == "fb" || g == "filebase" || g == "fn" || g == "filename" {
							core.FileBase(vallll, &tmp)
						} else if strings.HasPrefix(g, "json") {
							_, values := parse.ReadSqBr(g)
							core.GetJsonField(vallll, values, &tmp)
						} else if strings.HasPrefix(g, "case") {
							_, values := parse.ReadSqBr(g)
							core.Cases(vallll, values, &tmp)
						} else if strings.HasPrefix(g, "encode") {
							_, values := parse.ReadSqBr(g)
							core.Encode(vallll, values, &tmp)
						} else {
							core.AnalyzeURLs(vallll, g, &tmp)
						}

						vallll = tmp
						tmp = nil
					}
					columnValues = append(columnValues, vallll...)
					continue
				}
			}

			checkParamAndYaml(p, &columnValues)
		}

		if !appendMode[columnNum] || columnNum == 0 {
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

	// elapsed := time.Since(start)
	// fmt.Fprintf(os.Stderr, "Binomial took %s", elapsed)

	core.VPrint(fmt.Sprintf("Total words generated: %d", total))
}
