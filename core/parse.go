package core

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/giteshnxtlvl/cook/parse"
)

// var params = make(map[string]string)

var (
	Help    = false
	Verbose = false
	// Min     = 0
	// showConfig       = false
	ConfigPath = ""
	// caseValue        = ""
	// updateCacheFiles = false
	UpperCase = false
	LowerCase = false
)

var columnCases = make(map[int]map[string]bool)

func UpdateCases(caseValue string, noOfColumns int) map[int]map[string]bool {
	caseValue = strings.ToUpper(caseValue)

	for i := 0; i < noOfColumns; i++ {
		columnCases[i] = make(map[string]bool)
	}

	//Column Wise Cases
	if strings.Contains(caseValue, ":") {
		for _, val := range strings.Split(caseValue, ",") {
			v := strings.SplitN(val, ":", 2)
			i, err := strconv.Atoi(v[0])
			if err != nil {
				log.Fatalf("Err: Invalid column index %s", v[0])
			}
			for _, j := range strings.Split(v[1], "") {
				columnCases[i][j] = true
			}
		}
	} else {
		//Global Cases
		all := false
		if caseValue == "A" {
			all = true
			caseValue = ""
		}

		if all || strings.Contains(caseValue, "C") {
			columnCases[0]["L"] = true
			for i := 1; i < noOfColumns; i++ {
				columnCases[i]["T"] = true
			}
			caseValue = strings.ReplaceAll(caseValue, "C", "")
		}

		if all || strings.Contains(caseValue, "U") {
			UpperCase = true
			caseValue = strings.ReplaceAll(caseValue, "U", "")
		}

		if all || strings.Contains(caseValue, "L") {
			LowerCase = true
			caseValue = strings.ReplaceAll(caseValue, "L", "")
		}

		if all || strings.Contains(caseValue, "T") {
			for i := 0; i < noOfColumns; i++ {
				columnCases[i]["T"] = true
			}
		}

	}

	return columnCases
}

//Checking for patterns/functions
func ParseFunc(value string, array *[]string) bool {

	if !(strings.Contains(value, "(") && strings.Contains(value, ")")) {
		return false
	}

	funcName, funcValues := parse.ReadCrBr(value)
	// fmt.Println(funcName)
	// fmt.Println(funcValues)

	fmt.Print("")

	if funcPatterns, exists := M["patterns"][funcName]; exists {

		funcArgs := strings.Split(funcValues[0], ",")
		funcDef := strings.Split(funcPatterns[0][1:len(funcPatterns[0])-1], ",")

		// fmt.Printf("Func Arg: %v", funcArgs)
		// fmt.Printf("\tFunc Def: %v", funcDef)

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

var InputFile = make(map[string]bool)

func ParseFile(param string, value string, array *[]string) bool {

	// Checking for file
	if InputFile[param] && !strings.Contains(value, ":") {
		tmp := make(map[string]bool)
		FileValues(value, tmp)
		for k := range tmp {
			*array = append(*array, k)
		}
		return true
	}

	if checkFileInYaml(value, array) {
		return true
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
				return true
			} else if _, err := os.Stat(test2); err == nil {
				FindRegex([]string{one}, test2, array)
				return true
			}
		}

		if strings.Count(value, ":") == 1 {
			if _, err := os.Stat(value); err == nil {
				tmp := make(map[string]bool)
				FileValues(value, tmp)
				for k := range tmp {
					*array = append(*array, k)
				}
				return true
			}
			t := strings.SplitN(value, ":", 2)
			file, reg := t[0], t[1]

			if strings.HasSuffix(file, ".txt") {
				FindRegex([]string{file}, reg, array)
				return true
			} else if files, exists := M["files"][file]; exists {
				FindRegex(files, reg, array)
				return true
			}
		}
	}
	return false
}

var pipe []string

func PipeInput(value string, array *[]string) bool {
	if value == "-" {
		sc := bufio.NewScanner(os.Stdin)
		if len(pipe) > 0 {
			*array = append(*array, pipe...)
		}
		for sc.Scan() {
			*array = append(*array, sc.Text())
			pipe = append(pipe, sc.Text())
		}
		return true
	}
	return false
}

func RawInput(value string, array *[]string) bool {
	if strings.HasPrefix(value, "`") && strings.HasSuffix(value, "`") {
		lv := len(value)
		*array = append(*array, []string{value[1 : lv-1]}...)
		return true
	}
	return false
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

func ParsePorts(ports []string, array *[]string) {

	for _, p := range ports {

		if strings.Contains(p, "-") {
			pRange := strings.Split(p, "-")

			from, err1 := strconv.Atoi(pRange[0])
			till, err2 := strconv.Atoi(pRange[1])

			if err1 != nil || err2 != nil || from > till {
				log.Printf("Err: Wrong Range :/ %d-%d", from, till)
			}

			for i := from; i <= till; i++ {
				*array = append(*array, strconv.Itoa(i))
			}
		} else {
			port, err := strconv.Atoi(p)
			if err != nil {
				log.Printf("Err: Is this port number -_-?? '%s'", p)
			}
			*array = append(*array, strconv.Itoa(port))
		}
	}
}
