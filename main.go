package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"cook/core"
	"cook/parse"

	"github.com/ffuf/pencode/pkg/pencode"
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
	showCol          = parse.B("-col")
	Min              = parse.I("-m")
	appendColumns    = parse.S("-a")
	showConfig       = parse.B("-conf")
	caseValue        = parse.S("-c")
	encodeValue      = parse.S("-e")
	updateCacheFiles = parse.B("-u")
	l337             = parse.I("-l")
)

// var (
// 	help             = parse.B("h", "help")
// 	verbose          = parse.B("v", "verbose")
// 	showCol          = parse.B("c", "col")
// 	Min              = parse.I("m", "min")
// 	appendColumns    = parse.S("a", "append")
// 	showConfig       = parse.B("cf", "config")
// 	caseValue        = parse.S("ca", "case")
// 	encodeValue      = parse.S("e", "encode")
// 	updateCacheFiles = parse.B("u", "update")
// 	l337             = parse.I("l", "leet")
// )

var params = make(map[string]string)
var leetValues = make(map[string][]string)
var appendMode = make(map[int]bool)
var encodeString = []string{}
var finalFunc = func(s string) {
	fmt.Println(s)
}

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

func analyseParams(params map[string]string) {
	for param, value := range params {
		// fmt.Println(params)
		if strings.HasSuffix(param, ":") {
			delete(params, param)
			param = strings.TrimSuffix(param, ":")
			core.InputFile[param] = true
			params[param] = value
		}
	}
}

func searchMode(cmds []string) {
	core.CookConfig()

	search := cmds[0]
	found := false

	for cat, vv := range core.M {
		for k, v := range vv {

			if strings.Contains(k, search) {
				fmt.Println()
				if cat == "files" {
					fmt.Println(strings.ReplaceAll(k, search, core.Green+search+core.Reset))
					for _, file := range v {
						fmt.Printf("\t%s\n", strings.ReplaceAll(file, search, core.Green+search+core.Reset))
					}

				} else {
					fmt.Printf("%s \n\t%v\n", k, v)
				}
				found = true
			}
		}
	}

	if !found {
		fmt.Println("Not Found: ", search)
	}
	os.Exit(0)
}

func addMode(cmds []string) {
}
func updateMode(cmds []string) {
}
func deleteMode(cmds []string) {
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

	if len(encodeValue) > 0 {
		encodeString = strings.Split(encodeValue, ",")
		finalFunc = encode
	}

	core.Verbose = verbose

	params = parse.UserDefinedFlags()
	analyseParams(params)

	pattern := parse.Args
	noOfColumns := len(pattern)

	if noOfColumns > 0 {
		if pattern[0] == "search" {
			searchMode(pattern[1:])
		} else if pattern[0] == "help" {
			core.HelpMode(pattern[1:])
		} else if pattern[0] == "add" {
			addMode(pattern[1:])
		} else if pattern[0] == "update" {
			updateMode(pattern[1:])
		} else if pattern[0] == "delete" {
			deleteMode(pattern[1:])
		}
	}

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

	if len(appendColumns) > 0 {
		columns := strings.Split(appendColumns, ",")
		for _, colNum := range columns {
			intValue, err := strconv.Atoi(colNum)
			if err != nil {
				log.Fatalf("Err: Column Value %s in not integer", colNum)
			}
			appendMode[intValue] = true
		}
	}

	if showCol {
		fmt.Fprintln(os.Stderr)
		for i, p := range pattern {
			fmt.Fprintf(os.Stderr, "Col %d: %s\n", i, p)
		}
		fmt.Fprintln(os.Stderr)
		os.Exit(0)
	}

	core.VPrint(fmt.Sprintf("Pattern: %v \n", pattern))

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

	core.VPrint(fmt.Sprintf("Total words generated: %d", total))
}
