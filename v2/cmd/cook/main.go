package main

import (
	"fmt"
	"log"
	"time"

	"github.com/glitchedgitz/cook/v2/pkg/cook"
	"github.com/glitchedgitz/cook/v2/pkg/methods"
	"github.com/glitchedgitz/cook/v2/pkg/parse"
)



func initiate() {
	parse.Help = banner
	cook.Verbose = verbose
	help = parse.Boolean("-h", "-help")
	verbose = parse.Boolean("-v", "-verbose")
	showCol = parse.Boolean("-c", "-col")
	min = parse.Integer("-min", "-min")
	methodParam = parse.String("-mc", "-methodcol")
	methodsForAll = parse.String("-m", "-method")
	appendParam = parse.String("-a", "-append")
	showConfig = parse.Boolean("-conf", "-config")
	parse.Parse()
}

func main() {
	initiate()

	if help {
		// terminate after showing help
		showHelp()
	}

	if showConfig {
		// terminate after showing config
		showConf()
	}

	cook.CookConfig()

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

	run()

	cook.VPrint(fmt.Sprintf("%-40s: %s", "Elapsed Time", time.Since(start)))
	cook.VPrint(fmt.Sprintf("%-40s: %d", "Total words generated", total))
}

func init() {
	log.SetFlags(0)
}
