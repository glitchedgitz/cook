package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseCommand(list []string, val string) ([]string, bool) {
	for i, l := range list {
		if l == val {
			return append(list[:i], list[i+1:]...), true
		}
	}
	return list, false
}

func parseCommandArg(list []string, val string) ([]string, string) {
	for i, l := range list {
		if l == val {
			return append(list[:i], list[i+2:]...), list[i+1]
		}
	}
	return list, ""
}

func parseIntRanges(p string) ([]string, bool) {
	val := []string{}
	if strings.HasPrefix(p, "[") && strings.HasSuffix(p, "]") && strings.Contains(p, "-") {
		p = strings.ReplaceAll(p, "[", "")
		p = strings.ReplaceAll(p, "]", "")
		numRange := strings.SplitN(p, "-", 2)

		start, err := strconv.Atoi(numRange[0])
		if err != nil {
			return val, false
		}

		stop, err := strconv.Atoi(numRange[1])
		if err != nil {
			return val, false
		}

		for start <= stop {
			val = append(val, strconv.Itoa(start))
			start++
		}

		return val, true //get all ranges
	}
	return val, false
}

var columnCases = make(map[int][]string)

//Parse Input
func parseInput(commands []string) {

	if len(commands) == 0 {
		fmt.Println(banner)
		os.Exit(0)
	}

	if valueInSlice(commands, "-h") {
		showHelp()
	}

	if valueInSlice(commands, "-config") {
		showConfig()
	}

	commands, verbose = parseCommand(commands, "-v")
	commands, caseValue := parseCommandArg(commands, "-case")
	commands, minValue := parseCommandArg(commands, "-min")

	last := len(commands) - 1
	pattern = strings.Split(commands[last], ":")
	noOfColumns := len(pattern)

	if minValue == "" {
		min = noOfColumns - 1
	} else {
		var err error
		min, err = strconv.Atoi(minValue)
		min -= 1
		if err != nil {
			panic(err)
		}
	}

	if caseValue != "" {
		caseValue = strings.ToUpper(caseValue)
		if !strings.Contains(caseValue, ":") {
			tmp := strings.Split(caseValue, "")

			//For Camel Case Only
			if strings.Contains(caseValue, "C") {
				columnCases[0] = append(columnCases[0], "L")
				for i := 1; i < noOfColumns; i++ {
					columnCases[i] = append(columnCases[i], "T")
				}
			}

			for i := 0; i < noOfColumns; i++ {
				columnCases[i] = append(columnCases[i], tmp...)
			}
		} else {
			for _, val := range strings.Split(caseValue, ",") {
				v := strings.SplitN(val, ":", 2)
				i, err := strconv.Atoi(v[0])
				if err != nil {
					panic(err)
				}
				columnCases[i] = strings.Split(v[1], "")
			}
		}
	}

	for i, cmd := range commands[:last] {

		if strings.HasPrefix(cmd, "-") {
			cmd = strings.Replace(cmd, "-", "", 1)
			value := commands[i+1]
			values := []string{}

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

					for _, p := range funcPatterns[1:] {
						var tmp = p
						for index, arg := range funcDef {
							tmp = strings.ReplaceAll(tmp, arg, funcArgs[index])
						}
						values = append(values, tmp)
					}
					params[cmd] = values
					continue
				}
			}

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
						values = findRegex(test1, three)
					} else if _, err := os.Stat(test2); err == nil {
						values = findRegex(one, test2)
					} else {
						values = strings.Split(value, ",")
					}

				} else if strings.Count(value, ":") == 1 {
					if _, err := os.Stat(value); err == nil {
						values = fileValues(value)
					} else {
						t := strings.SplitN(value, ":", 2)
						file := t[0]
						reg := t[1]

						if strings.HasSuffix(file, ".txt") {
							values = findRegex(file, reg)
						} else if _, exists := m["files"][file]; exists {
							values = findRegex(m["files"][file][0], reg)
						} else {
							values = strings.Split(value, ",")
						}
					}
				}

			} else if strings.HasSuffix(value, ".txt") {
				if _, err := os.Stat(value); err == nil {
					values = fileValues(value)
				}
			} else {
				values = strings.Split(value, ",")
			}

			params[cmd] = values
		}
	}
}

func showHelp() {
	fmt.Println(banner)

	fmt.Println(green + "\nGITHUB" + white)
	fmt.Println(blue + "    https://github.com/giteshnxtlvl/cook" + reset)

	fmt.Println(green + "\nFLAGS" + white)
	help := `    -case   : Define Cases
              * Use for complete list
                  -case A for ALL 
                  -case U for Uppercase
                  -case L for Lowercase
                  -case T for Titlecase
                  -case C for Camelcase

              * Use column wise, no camel case for this
                  -case 0:U,2:T
                      Column 0 will be in Uppercase
                      Column 2 will be in Titlecase,
                      Rest columns will be default output
                  Multiple Cases
                      -case 0:UT,2:A 

    -min    : Minimum no of columns to print. (Default min = no of columns)
              Same as minimum of crunch			  
    -config : Config Information *cook.yaml*
    -v      : Verbose Output 
    -h      : Help
	`
	fmt.Println(help)

	fmt.Println(green + "\nBASIC USAGE" + white)
	fmt.Printf("   $ cook %[1]s-start %[2]sadmin%[3]s,%[2]sroot  %[1]s-sep %[2]s_%[3]s,%[2]s-  %[1]s-end %[2]ssecret%[3]s,%[2]scritical  %[2]s/%[3]s:%[1]sstart%[3]s:%[1]ssep%[3]s:%[1]send\n", green, blue, white)
	fmt.Printf("   %[3]s$ cook %[2]s/%[3]s:%[2]sadmin%[3]s,%[2]sroot%[3]s:%[2]s_%[3]s,%[2]s-%[3]s:%[2]ssecret%[3]s,%[2]scritical\n", green, blue, white)

	fmt.Println(green + "\nFILE WITH REGEX" + white)
	fmt.Printf("   $ cook %[1]s-s %[2]scompany %[1]s-ext %[2]sraft-large-extensions%[3]s:%[3]s\\.asp.*  %[2]s/%[3]s:%[1]ss%[3]s:%[1]sexp\n", green, blue, white, purple)

	os.Exit(0)
}

func showConfig() {

	configFile := os.Getenv("COOK")
	fmt.Println(green + "\nCOOK.YAML " + reset)
	if len(configFile) == 0 {
		fmt.Println("  You are using Default Sets" + reset)
	} else {
		fmt.Println("  You are using Custom cook.yaml" + reset)
		fmt.Printf(blue+"  %-12s "+white+": %v\n", "Location", configFile)
	}

	fmt.Println(green + "\nCHARACTER SETS" + reset)
	for k, v := range m["charSet"] {
		fmt.Printf(blue+"  %-12s "+white+"%v\n", k, v[0])
	}
	fmt.Println(green + "\nFILES" + reset)
	for k, v := range m["files"] {
		fmt.Printf(blue+"  %-12s "+white+"%s\n", k, v[0])
	}
	fmt.Println(green + "\nLISTS" + reset)
	for k, v := range m["lists"] {
		fmt.Printf(blue+"  %-12s "+white+"%v\n", k, v)
	}
	fmt.Println(green + "\nPATTERNS" + reset)
	for k, v := range m["patterns"] {
		fmt.Printf(blue+"  %-12s "+white+"%v\n", k, v)
	}
	fmt.Println(green + "\nEXTENSIONS" + reset)
	for k, v := range m["extensions"] {
		fmt.Printf(blue+"  %-12s "+white+"%v\n", k, v)
	}
	os.Exit(0)
}
