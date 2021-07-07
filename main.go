package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/giteshnxtlvl/cook/core"
	"github.com/giteshnxtlvl/cook/parse"

	"github.com/giteshnxtlvl/pencode/pkg/pencode"
)

var total = 0
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

var appendMode = make(map[int]bool)
var encodeString []string

var finalFunc = func(s string) {
	fmt.Println(s)
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

func checkParam(p string, array *[]string) bool {
	if val, exists := params[p]; exists {
		if core.PipeInput(val, array) || core.RawInput(val, array) || core.ParseFunc(val, array) || core.ParseFile(p, val, array) || checkMethods(val, array) {
			return true
		}

		*array = append(*array, strings.Split(val, ",")...)
		return true
	}
	return false
}

func init() {
	log.SetFlags(0)
}

func checkMethods(p string, array *[]string) bool {
	if strings.Count(p, ".") > 0 {
		splitS := strings.Split(p, ".")
		u := splitS[0]
		if _, exists := params[u]; exists {

			get := splitS[1:]

			tmp := []string{}
			vallll := []string{}

			if !checkParam(u, &vallll) && !core.CheckYaml(u, &vallll) {
				return false
			}

			vallll = append(vallll, p)

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
			*array = append(*array, vallll...)
			return true
		}
	}
	return false
}

var params map[string]string
var pattern []string

func main() {
	start := time.Now()
	params, pattern = parseInput()

	core.CookConfig()

	for columnNum, param := range pattern {

		columnValues := []string{}

		for _, p := range strings.Split(param, ",") {
			core.VPrint(fmt.Sprintf("Param: %s \n", p))

			if core.RawInput(p, &columnValues) || core.ParseRanges(p, &columnValues) || core.ParseFunc(p, &columnValues) || checkMethods(p, &columnValues) || checkParam(p, &columnValues) || core.CheckYaml(p, &columnValues) {
				continue
			}
			columnValues = append(columnValues, p)
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

	core.VPrint(fmt.Sprintf("Elapsed Time: %s", time.Since(start)))
	core.VPrint(fmt.Sprintf("Total words generated: %d", total))
}
