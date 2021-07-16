package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/giteshnxtlvl/cook/pkg/cook"
	"github.com/giteshnxtlvl/cook/pkg/parse"
)

var home, _ = os.UserHomeDir()
var totalCols = 0

var (
	help          = parse.B("-h", "-help")
	verbose       = parse.B("-v", "-verbose")
	showCol       = parse.B("-c", "-col")
	min           = parse.I("-m", "-min")
	appendColumns = parse.S("-a", "-append")
	showConfig    = parse.B("-conf", "-config")
	caseValue     = parse.S("-ca", "-case")
	encodeValue   = parse.S("-e", "-encode")
	l337          = parse.I("-l", "-leet")
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

	if len(encodeValue) > 0 {
		encodeString = strings.Split(encodeValue, ":")
		doEncode = true
	}

	params = parse.UserDefinedFlags()
	pattern = parse.Args

	totalCols = len(pattern)

	analyseParams(params)

	cmdsMode()
	setMin()

	if caseValue != "" {
		columnCases = cook.UpdateCases(caseValue, totalCols)
	}

	if l337 > -1 {
		doLeet = true
		if l337 > 1 {
			fmt.Println("Err: -1337 can be 0 or 1, 0 - Calm Mode & 1 - Angry Mode", l337)
			os.Exit(0)
		}
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
		showCols()
	}

	cook.VPrint(fmt.Sprintf("Pattern: %v \n", pattern))

}
