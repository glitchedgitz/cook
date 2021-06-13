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
			vPrint(fmt.Sprintf("Param: %s Value: %s", cmd, value))
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

func parseRanges(p string, array *[]string) bool {

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
				*array = append(*array, strconv.Itoa(start))
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
					*array = append(*array, charsList[start])
					start++
				}
				success = true
			}
		}
	}
	return success
}

var configPath string
var commands []string
var verbose = false

func vPrint(msg string) {
	if verbose {
		fmt.Fprintln(os.Stderr, msg)
	}
}

func parseInput() {

	if len(commands) == 0 {
		fmt.Println(banner)
		os.Exit(0)
	}

	if parseBoolArg("-h") {
		showHelp()
	}

	configPath = parseStringArg("-config-path")

	if parseBoolArg("-config") {
		cookConfig()
		showConfig()
	}

	if parseBoolArg("-update-all") {
		cookConfig()
		updateCache()
		os.Exit(0)
	}

	verbose = parseBoolArg("-v")
	caseValue := parseStringArg("-case")
	min = parseIntArg("-min")

	tmp := []string{}

	tmp = append(tmp, commands...)

	for _, cmd := range tmp {
		if len(cmd) > 1 && strings.HasPrefix(cmd, "-") {
			value := parseStringArg(cmd)
			cmd = strings.Replace(cmd, "-", "", 1)
			params[cmd] = value
		}
	}

	pattern = commands
	noOfColumns := len(pattern)

	if min == 0 {
		min = noOfColumns - 1
	} else {
		min -= 1
	}

	if caseValue != "" {
		updateCases(caseValue, noOfColumns)
	}

}

//Checking for patterns/functions
func parseFunc(value string, array *[]string) bool {

	function := strings.SplitN(value, "(", 2)
	funcName := function[0]

	if funcPatterns, exists := m["patterns"][funcName]; exists {

		funcArgs := strings.Split(function[1][:len(function[1])-1], ",")
		funcDef := strings.Split(strings.TrimSuffix(funcPatterns[0][len(funcName)+1:], ")"), ",")

		if len(funcDef) != len(funcArgs) {
			log.Fatalf(red+"\nErr: No of Arguments are different for %s\n", funcPatterns)
		}

		for _, p := range funcPatterns[1:] {
			for index, arg := range funcDef {
				p = strings.ReplaceAll(p, arg, funcArgs[index])
			}
			*array = append(*array, p)
		}

		return true
	}
	return false
}

func parseValue(value string, array *[]string) {

	// Pipe input
	if value == "-" {
		sc := bufio.NewScanner(os.Stdin)

		for sc.Scan() {
			*array = append(*array, sc.Text())
		}
		return
	}

	if strings.HasPrefix(value, "`") && strings.HasSuffix(value, "`") {
		lv := len(value)
		*array = append(*array, []string{value[1 : lv-1]}...)
		return
	}

	success := parseFunc(value, array)
	if success {
		return
	}

	// Checking for file
	if strings.HasSuffix(value, ".txt") {
		if _, err := os.Stat(value); err == nil {
			fileValues(value, array)
			return
		}
	}

	// Checking for File and Regex
	if strings.Contains(value, ":") {
		// File may starts from E: C: D: for windows + Regex is supplied
		if strings.Count(value, ":") == 2 {
			tmp := strings.SplitN(value, ":", 3)

			one, two, three := tmp[0], tmp[1], tmp[2]
			test1, test2 := one+":"+two, two+":"+three

			if _, err := os.Stat(test1); err == nil {
				findRegex([]string{test1}, three, array)
				return
			} else if _, err := os.Stat(test2); err == nil {
				findRegex([]string{one}, test2, array)
				return
			}
		}

		if strings.Count(value, ":") == 1 {
			if _, err := os.Stat(value); err == nil {
				fileValues(value, array)
				return
			}
			t := strings.SplitN(value, ":", 2)
			file, reg := t[0], t[1]

			if strings.HasSuffix(file, ".txt") {
				findRegex([]string{file}, reg, array)
				return
			} else if files, exists := m["files"][file]; exists {
				findRegex(files, reg, array)
				return
			}
		}
	}

	*array = append(*array, strings.Split(value, ",")...)
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
				log.Fatalln("Err: Invalid column index for cases")
			}
			for _, j := range strings.Split(v[1], "") {
				columnCases[i][j] = true
			}
		}
	}
}
