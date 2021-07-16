package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/giteshnxtlvl/cook/pkg/cook"
	"github.com/giteshnxtlvl/cook/pkg/parse"
)

var total = 0

// var otherCases = false
var columnCases = make(map[int]map[string]bool)

//Initializing with empty string, so loops will run for 1st column
var final = []string{""}

// -save [keyword] to save the generated permutations
var appendMode = make(map[int]bool)
var encodeString []string

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
	tmp := []string{}

	if len(columnCases[columnNum]) > 0 {
		// otherCases = true
		allcases := false

		if columnCases[columnNum]["A"] {
			allcases = true
		}

		if allcases || (!cook.UpperCase && columnCases[columnNum]["U"]) {
			applyFunc(columnValues, &tmp, strings.ToUpper)
		}

		if allcases || columnCases[columnNum]["L"] {
			applyFunc(columnValues, &tmp, strings.ToLower)
		}

		if allcases || columnCases[columnNum]["T"] {
			applyFunc(columnValues, &tmp, strings.Title)
		}

	} else {
		applyFunc(columnValues, &tmp, useless)
	}

	final = tmp
}

func checkParam(p string, array *[]string) bool {
	if val, exists := params[p]; exists {
		if cook.PipeInput(val, array) || cook.RawInput(val, array) || cook.ParseFunc(val, array) || cook.ParseFile(p, val, array) || checkMethods(val, array) {
			return true
		}

		*array = append(*array, splitValues(val)...)
		return true
	}
	return false
}

func splitMethods(p string) []string {
	chars := strings.Split(p, "")
	s := []string{}
	tmp := ""
	insidebrackets := false
	for _, c := range chars {

		if c == "." {
			if !insidebrackets {
				s = append(s, tmp)
				tmp = ""
				continue
			}
		}
		if c == "[" {
			insidebrackets = true
		}
		if c == "]" {
			insidebrackets = false
		}
		tmp += c
	}
	s = append(s, tmp)
	return s
}

func checkMethods(p string, array *[]string) bool {
	if strings.Count(p, ".") > 0 {
		splitS := splitMethods(p)
		u := splitS[0]
		if _, exists := params[u]; exists {

			get := splitS[1:]

			tmp := []string{}
			vallll := []string{}

			if !checkParam(u, &vallll) && !cook.CheckYaml(u, &vallll) {
				return false
			}

			// vallll = append(vallll, p)

			for _, g := range get {
				if g == "smart" {
					cook.SmartWords(vallll, useless, &tmp)

				} else if strings.HasPrefix(g, "smartjoin") {
					_, values := parse.ReadSqBrSepBy(g, ":")
					cook.SmartWordsJoin(vallll, values, useless, &tmp)

				} else if g == "fb" || g == "filebase" || g == "fn" || g == "filename" {
					cook.FileBase(vallll, &tmp)

				} else if strings.HasPrefix(g, "regex") {
					_, value := parse.ReadSqBr(g)
					cook.Regex(vallll, value, &tmp)

				} else if strings.HasPrefix(g, "split") {
					_, value := parse.ReadSqBr(g)
					cook.Split(vallll, value, &tmp)

				} else if strings.HasPrefix(g, "splitindex") {
					_, values := parse.ReadSqBrSepBy(g, ":")
					index, err := strconv.Atoi(values[1])

					if err != nil {
						log.Fatalln("Not Integer Value: "+values[1], err)
					}

					cook.SplitIndex(vallll, values[0], index, &tmp)

				} else if strings.HasPrefix(g, "json") {

					_, values := parse.ReadSqBrSepBy(g, ":")
					cook.GetJsonField(vallll, values, &tmp)

				} else if strings.HasPrefix(g, "case") {

					_, values := parse.ReadSqBrSepBy(g, ":")
					cook.Cases(vallll, values, &tmp)

				} else if strings.HasPrefix(g, "leet") {

					_, value := parse.ReadSqBr(g)
					mode, err := strconv.Atoi(value)
					if err != nil {
						log.Fatalln("Err: Leet can be 0 or 1")
					}
					cook.Leet(vallll, mode, &tmp)

				} else if strings.HasPrefix(g, "encode") {

					_, values := parse.ReadSqBrSepBy(g, ":")
					cook.Encode(vallll, values, &tmp)

				} else {

					cook.AnalyzeURLs(vallll, g, &tmp)

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

var start = time.Now()

func splitValues(p string) []string {
	chars := strings.Split(p, "")
	s := []string{}
	tmp := ""
	insideraw := false

	for _, c := range chars {

		if c == "," {
			if !insideraw {
				s = append(s, tmp)
				tmp = ""
				continue
			}
		}
		if c == "`" {
			if insideraw {
				insideraw = false
			} else {
				insideraw = true
			}
		}

		tmp += c
	}
	s = append(s, tmp)
	return s
}

func main() {

	parseInput()

	for columnNum, param := range pattern {

		columnValues := []string{}

		for _, p := range splitValues(param) {
			cook.VPrint(fmt.Sprintf("Param: %s \n", p))
			if cook.RawInput(p, &columnValues) || cook.ParseRanges(p, &columnValues) || checkMethods(p, &columnValues) || checkParam(p, &columnValues) || cook.CheckYaml(p, &columnValues) {
				continue
			}
			columnValues = append(columnValues, p)

		}

		cook.VPrint(fmt.Sprintf("%-40s: %s", "Time after getting values", time.Since(start)))

		if !appendMode[columnNum] || columnNum == 0 {
			applyColumnCases(columnValues, columnNum, applyCase)
		} else {
			applyColumnCases(columnValues, columnNum, prefixSufixMode)
		}

		cook.VPrint(fmt.Sprintf("%-40s: %s", "Time ApplyColumnCases", time.Since(start)))

		if columnNum >= min {
			print()
		}
	}

	cook.VPrint(fmt.Sprintf("%-40s: %s", "Elapsed Time", time.Since(start)))
	cook.VPrint(fmt.Sprintf("%-40s: %d", "Total words generated", total))
}

func init() {
	log.SetFlags(0)
}
