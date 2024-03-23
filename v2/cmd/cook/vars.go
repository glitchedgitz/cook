package main

import (
	"os"
	"time"
)

var total = 0

// Initializing with empty string, so loops will run for 1st column
var final = []string{""}
var params map[string]string
var pattern []string
var start = time.Now()

var home, _ = os.UserHomeDir()
var totalCols = 0
var methodMap = make(map[int][]string)

// Flags
var (
	help          bool
	verbose       bool
	showCol       bool
	min           int
	methodParam   string
	methodsForAll string
	appendParam   string
	showConfig    bool
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
	"meths":   methHelp,
	"usage":   usageHelp,
	"flags":   flagsHelp,
}
