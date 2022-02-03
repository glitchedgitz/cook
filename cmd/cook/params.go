package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/giteshnxtlvl/cook/v2/pkg/cook"
	"github.com/giteshnxtlvl/cook/v2/pkg/methods"
	"github.com/giteshnxtlvl/cook/v2/pkg/parse"
)

var home, _ = os.UserHomeDir()
var totalCols = 0

var (
	help          = parse.B("-h", "-help")
	verbose       = parse.B("-v", "-verbose")
	showCol       = parse.B("-c", "-col")
	min           = parse.I("-min", "-min")
	methodParam   = parse.S("-mc", "-methodcol")
	methodsForAll = parse.S("-m", "-method")
	appendParam   = parse.S("-a", "-append")
	showConfig    = parse.B("-conf", "-config")
)

var cmdFunctions = map[string]func([]string){
	"search": searchMode,
	"show":   showMode,
	"help":   helpMode,
	"add":    addMode,
	"clean":  cleanMode,
	"info":   infoMode,
	"update": updateMode,
	"delete": deleteMode,
}

func showConf() {
	fmt.Println(cook.Blue + "\n    CONFIG" + cook.Reset)
	fmt.Printf("    Location: %v\n", cook.ConfigFolder)
	fmt.Printf(cook.Blue+"\n    %-25s   %s     %s   %s   %s\n"+cook.Reset, "FILE", "SETS", "VERN", "PREFIX", "REPO")
	fmt.Println(cook.ConfigInfo)

	os.Exit(0)
}

// Todo: replace this func whereevery needed
func getInt(a string) int {
	num, err := strconv.Atoi(a)
	if err != nil {
		log.Fatalf("Err: \"%s\" is not integer", a)
	}
	return num
}

func setMin() {
	if min < 0 {
		min = totalCols - 1
	} else {
		if min > totalCols {
			fmt.Println("Err: min is greator than no of columns")
			os.Exit(0)
		}
		min -= 1
	}
}

func showCols() {
	fmt.Fprintln(os.Stderr)
	for i, p := range pattern {
		fmt.Fprintf(os.Stderr, "Col %d: %s\n", i, p)
	}
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func cmdsMode() {
	if totalCols > 0 {
		if fn, exists := cmdFunctions[pattern[0]]; exists {
			fn(pattern[1:])
			os.Exit(0)
		}
	}
}

var appendMap = make(map[int]bool)

func parseAppend() {
	columns := strings.Split(appendParam, ",")
	for _, colNum := range columns {
		intValue, err := strconv.Atoi(colNum)
		if err != nil {
			log.Fatalf("Err: Column Value %s in not integer", colNum)
		}
		appendMap[intValue] = true
	}
}

var methodMap = make(map[int][]string)

func parseMethod() {
	meths := strings.Split(methodParam, ";")
	forAllCols := []string{}

	var modifiedCol = make(map[int]bool)

	for _, m := range meths {
		if strings.Contains(m, ":") {
			s := strings.SplitN(m, ":", 2)
			i := getInt(s[0])
			if i >= totalCols {
				log.Fatalf("Err: No Column %d", i)
			}
			methodMap[i] = strings.Split(s[1], ",")
			modifiedCol[i] = true
		} else {
			forAllCols = append(forAllCols, strings.Split(m, ",")...)
		}
	}

	for i := 0; i < totalCols; i++ {
		if !modifiedCol[i] {
			methodMap[i] = forAllCols
		}
	}
}

func parseInput() {
	parse.Help = banner
	cook.Verbose = verbose

	parse.Parse()

	if help {
		showHelp()
	}

	cook.CookConfig()

	if showConfig {
		showConf()
	}

	params = parse.UserDefinedFlags()
	pattern = parse.Args

	totalCols = len(pattern)

	analyseParams(params)
	cmdsMode()
	setMin()
	methods.LeetBegin()

	if len(appendParam) > 0 {
		parseAppend()
	}

	if len(methodParam) > 0 {
		parseMethod()
	}

	if showCol {
		showCols()
	}

	cook.VPrint(fmt.Sprintf("Pattern: %v \n", pattern))
}
