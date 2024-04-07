package main

import (
	"fmt"
	"os"

	"github.com/glitchedgitz/cook/v2/pkg/util"
)

// cook [cmdFunctions]
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

// cook help [helpFunctions]
var helpFunctions = map[string]func(){
	"methods": methHelp,
	"usage":   usageHelp,
	"flags":   flagsHelp,
}

func showConf() {
	// cook.CookConfig()

	fmt.Println(util.Blue + "\n    CONFIG" + util.Reset)
	fmt.Printf("    Location: %v\n", COOK.Config.ConfigPath)
	fmt.Printf(util.Blue+"\n    %-25s   %s     %s   %s   %s\n"+util.Reset, "FILE", "SETS", "VERN", "PREFIX", "REPO")
	fmt.Println(COOK.Config.ConfigInfo)

	os.Exit(0)
}
