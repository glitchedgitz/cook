package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseBoolArg(flag string) bool {
	for i, cmd := range commands {
		if cmd == flag {
			commands = append(commands[:i], commands[i+1:]...)
			return true
		}
	}
	return false
}

func parseStringArg(flag string) string {
	for i, cmd := range commands {
		if cmd == flag {
			value := commands[i+1]
			commands = append(commands[:i], commands[i+2:]...)
			return value
		}
	}

	return ""
}

func parseIntArg(flag string) int {
	intValue := 0
	for i, l := range commands {
		if l == flag {
			if commands[i+1] == "" {
				log.Fatalf("Err: Flag %s don't have value", flag)
				// min = noOfColumns - 1
			} else {
				var err error
				intValue, err = strconv.Atoi(commands[i+1])
				// min -= 1
				if err != nil {
					log.Fatalf("Err: Flag %s needs integer value", flag)
				}
			}
			commands = append(commands[:i], commands[i+2:]...)
			return intValue
		}
	}
	return 0
}

func parseRanges(p string) ([]string, bool) {
	val := []string{}
	success := false
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	if strings.HasPrefix(p, "[") && strings.HasSuffix(p, "]") && strings.Contains(p, "-") {

		p = strings.ReplaceAll(strings.ReplaceAll(p, "[", ""), "]", "")
		numRange := strings.SplitN(p, "-", 2)
		from := numRange[0]
		to := numRange[1]

		start, err1 := strconv.Atoi(from)
		stop, err2 := strconv.Atoi(to)

		if err1 == nil && err2 == nil {
			for start <= stop {
				val = append(val, strconv.Itoa(start))
				start++
			}
			success = true
		}

		if !success && len(from) == 1 && len(to) == 1 && strings.Contains(chars, from) && strings.Contains(chars, to) {
			start = strings.Index(chars, from)
			stop = strings.Index(chars, to)

			if start < stop {
				charsList := strings.Split(chars, "")
				for start <= stop {
					val = append(val, charsList[start])
					start++
				}
				success = true
			}
		}
	}
	return val, success
}

var configPath string
var commands []string
var verbose bool

func vPrint(msg string) {
	if verbose {
		fmt.Fprintln(os.Stderr, msg)
	}
}

func parseInput() {

	if len(commands) == 0 {
		os.Exit(0)
	}

	last := len(commands) - 1
	pattern = strings.Split(commands[last], ":")
	noOfColumns := len(pattern)

	verbose = parseBoolArg("-v")
	if parseBoolArg("-h") {
		showHelp()
	}
	configPath = parseStringArg("-config-path")

	if parseBoolArg("-config") {
		cookConfig()
		showConfig()
	}

	caseValue := parseStringArg("-case")
	min = parseIntArg("-min")

	if min == 0 {
		min = noOfColumns - 1
	} else {
		min -= 1
	}

	if caseValue != "" {
		updateCases(caseValue, noOfColumns)
	}

	for i, cmd := range commands[:last] {
		if strings.HasPrefix(cmd, "-") {
			cmd = strings.Replace(cmd, "-", "", 1)
			value := commands[i+1]
			vPrint(fmt.Sprintf("%12s %s", cmd, value))
			params[cmd] = value
		}
	}
}

func parseValue(value string) []string {

	// Pipe input
	if value == "-" {
		tmp := []string{}
		sc := bufio.NewScanner(os.Stdin)

		for sc.Scan() {
			tmp = append(tmp, sc.Text())
		}
		return tmp
	}

	// Raw String using `
	if strings.HasPrefix(value, "`") && strings.HasSuffix(value, "`") {
		return []string{value}
	}

	//Checking for patterns/functions
	if strings.Contains(value, "(") && strings.HasSuffix(value, ")") {
		function := strings.SplitN(value, "(", 2)
		funcName := function[0]
		if _, exists := m["patterns"][funcName]; exists {
			funcArgs := strings.Split(strings.TrimSuffix(function[1], ")"), ",")
			funcPatterns := m["patterns"][funcName]
			funcDef := strings.Split(strings.TrimSuffix(strings.TrimPrefix(funcPatterns[0], funcName+"("), ")"), ",")

			if len(funcDef) != len(funcArgs) {
				fmt.Printf(red+"\nError: No of Arguments are different for %s\n", funcPatterns[0])
				os.Exit(0)
			}

			values := []string{}
			for _, p := range funcPatterns[1:] {
				var tmp = p
				for index, arg := range funcDef {
					tmp = strings.ReplaceAll(tmp, arg, funcArgs[index])
				}
				values = append(values, tmp)
			}
			return values
		}
	}

	// Checking for File and Regex
	if strings.Contains(value, ":") {
		if strings.Count(value, ":") == 2 {
			// File is starting from E: C: D: for windows + Regex is supplied
			tmp := strings.SplitN(value, ":", 3)

			one := tmp[0]
			two := tmp[1]
			three := tmp[2]
			test1 := one + ":" + two
			test2 := two + ":" + three

			if _, err := os.Stat(test1); err == nil {
				return findRegex(test1, three)
			} else if _, err := os.Stat(test2); err == nil {
				return findRegex(one, test2)
			} else {
				return strings.Split(value, ",")
			}

		} else if strings.Count(value, ":") == 1 {
			if _, err := os.Stat(value); err == nil {
				return fileValues(value)
			}
			t := strings.SplitN(value, ":", 2)
			file := t[0]
			reg := t[1]

			if strings.HasSuffix(file, ".txt") {
				return findRegex(file, reg)
			} else if _, exists := m["files"][file]; exists {
				return findRegex(m["files"][file][0], reg)
			} else {
				return strings.Split(value, ",")
			}
		}
	} else if strings.HasSuffix(value, ".txt") {
		if _, err := os.Stat(value); err == nil {
			return fileValues(value)
		}
	}

	return strings.Split(value, ",")
}

var columnCases = make(map[int]map[string]bool)

func updateCases(caseValue string, noOfColumns int) {
	caseValue = strings.ToUpper(caseValue)

	for i := 0; i < noOfColumns; i++ {
		columnCases[i] = make(map[string]bool)
	}

	//Global Cases
	if !strings.Contains(caseValue, ":") {

		//For Camel Case Only
		if strings.Contains(caseValue, "C") {
			columnCases[0]["L"] = true
			for i := 1; i < noOfColumns; i++ {
				columnCases[i]["T"] = true
			}
		}

		for i := 0; i < noOfColumns; i++ {
			for _, c := range strings.Split(caseValue, "") {
				columnCases[i][c] = true
			}
		}
	} else { //Column Wise Cases
		for _, val := range strings.Split(caseValue, ",") {
			v := strings.SplitN(val, ":", 2)
			i, err := strconv.Atoi(v[0])
			if err != nil {
				fmt.Println("Err: Invalid column index for cases")
			}
			for _, j := range strings.Split(v[1], "") {
				columnCases[i][j] = true
			}
		}
	}
}
