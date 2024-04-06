package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/glitchedgitz/cook/v2/pkg/config"
	"github.com/glitchedgitz/cook/v2/pkg/cook"
	"github.com/glitchedgitz/cook/v2/pkg/parse"
)

var COOK *cook.COOK
var configPath string
var total = 0

// Initializing with empty string, so loops will run for 1st column
var start = time.Now()

// Flags
var (
	help          bool
	verbose       bool
	showCol       bool
	min           int
	peek          int
	methodParam   string
	methodsForAll string
	appendParam   string
	showConfig    bool
	reConfigure   bool
)

func main() {
	parseFlags := parse.NewParse()
	parseFlags.Help = banner
	// cook.Verbose = verbose
	help = parseFlags.Boolean("-h", "-help")
	verbose = parseFlags.Boolean("-v", "-verbose")
	showCol = parseFlags.Boolean("-c", "-col")
	min = parseFlags.Integer("-min", "-min")
	peek = parseFlags.Integer("-peek", "-peek")
	methodParam = parseFlags.String("-mc", "-methodcol")
	methodsForAll = parseFlags.String("-m", "-method")
	appendParam = parseFlags.String("-a", "-append")
	showConfig = parseFlags.Boolean("-conf", "-config")
	reConfigure = parseFlags.Boolean("-reconfigure", "-reconf")
	configPath = parseFlags.String("-config-path", "-config-path")
	parseFlags.Parse()

	if !verbose {
		log.SetFlags(0)
	}

	if help {
		// terminate after showing help
		showHelp()
	}

	if showConfig {
		// terminate after showing config
		showConf()
	}

	if len(os.Getenv("COOK")) > 0 {
		configPath = os.Getenv("COOK")
	}

	COOK = cook.New(&cook.COOK{
		Config: &config.Config{
			ConfigPath:  configPath,
			ReConfigure: reConfigure,
			Verbose:     verbose,
			Peek:        peek,
		},
		Pattern:       parseFlags.Args,
		Min:           min,
		AppendParam:   appendParam,
		MethodParam:   methodParam,
		MethodsForAll: methodsForAll,
		PrintResult:   true,
	})

	if showCol {
		COOK.ShowCols()
	}

	if COOK.TotalCols > 0 {
		if fn, exists := cmdFunctions[COOK.Pattern[0]]; exists {
			fn(COOK.Pattern[1:])
			os.Exit(0)
		}
	}

	VPrint(fmt.Sprintf("Pattern: %v \n", COOK.Pattern))
	// COOK.CurrentStage()
	COOK.Generate()

	if verbose {
		COOK.CurrentStage()
	}

	VPrint(fmt.Sprintf("%-40s: %s", "Elapsed Time", time.Since(start)))
	VPrint(fmt.Sprintf("%-40s: %d", "Total words generated", total))
}
