package core

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"cook/pkg/parse"
)

var params = make(map[string]string)

var (
	help             = parse.Bool("-h")
	verbose          = parse.Bool("-v")
	Min              = parse.Int("-Min")
	showConfig       = parse.Bool("-config")
	configPath       = parse.String("-config-path")
	caseValue        = parse.String("-case")
	updateCacheFiles = parse.Bool("-update-all")
)

func ParseInput() (map[string]string, []string) {

	if help {
		showHelp()
	}

	if showConfig {
		CookConfig()
		ShowConfig()
	}

	if updateCacheFiles {
		CookConfig()
		UpdateCache()
		os.Exit(0)
	}

	params = parse.UserDefinedFlags()

	pattern := parse.Commands
	noOfColumns := len(pattern)

	if Min == 0 {
		Min = noOfColumns - 1
	} else {
		Min -= 1
	}

	if caseValue != "" {
		updateCases(caseValue, noOfColumns)
	}

	return params, pattern
}

//Checking for patterns/functions
func ParseFunc(value string, array *[]string) bool {

	function := strings.SplitN(value, "(", 2)
	funcName := function[0]

	if funcPatterns, exists := M["patterns"][funcName]; exists {

		funcArgs := strings.Split(function[1][:len(function[1])-1], ",")
		funcDef := strings.Split(strings.TrimSuffix(funcPatterns[0][len(funcName)+1:], ")"), ",")

		if len(funcDef) != len(funcArgs) {
			log.Fatalf("\nErr: No of Arguments are different for %s\n", funcPatterns)
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

func ParseValue(value string, array *[]string) {

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

	success := ParseFunc(value, array)
	if success {
		return
	}

	// Checking for file
	if strings.HasSuffix(value, ".txt") {
		if _, err := os.Stat(value); err == nil {
			FileValues(value, array)
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
				FindRegex([]string{test1}, three, array)
				return
			} else if _, err := os.Stat(test2); err == nil {
				FindRegex([]string{one}, test2, array)
				return
			}
		}

		if strings.Count(value, ":") == 1 {
			if _, err := os.Stat(value); err == nil {
				FileValues(value, array)
				return
			}
			t := strings.SplitN(value, ":", 2)
			file, reg := t[0], t[1]

			if strings.HasSuffix(file, ".txt") {
				FindRegex([]string{file}, reg, array)
				return
			} else if files, exists := M["files"][file]; exists {
				FindRegex(files, reg, array)
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

func ParseRanges(p string, array *[]string) bool {

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
