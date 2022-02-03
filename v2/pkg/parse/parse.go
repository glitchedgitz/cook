package parse

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var Args = os.Args[1:]
var showError = false
var Help = ""

func B(flag, flagL string) bool {
	for i, cmd := range Args {
		if cmd == flag || cmd == flagL {
			Args = append(Args[:i], Args[i+1:]...)
			return true
		}
	}
	return false
}

func S(flag, flagL string) string {
	l := len(Args)
	for i, cmd := range Args {
		if cmd == flag || cmd == flagL {
			if i+1 == l {
				fmt.Printf("Err: Flag '%s' doesn't have any value", cmd)
				os.Exit(0)
			}
			value := Args[i+1]
			Args = append(Args[:i], Args[i+2:]...)
			return value
		}
	}

	return ""
}

func I(flag, flagL string) int {
	intValue := 0
	l := len(Args)

	for i, cmd := range Args {
		if cmd == flag || cmd == flagL {
			if i+1 == l || Args[i+1] == "" {
				fmt.Printf("Err: Flag '%s' doesn't have any value", cmd)
				os.Exit(0)
			} else {
				var err error
				intValue, err = strconv.Atoi(Args[i+1])
				// min -= 1
				if err != nil {
					log.Fatalf("Err: Flag %s needs integer value", flag)
				}
			}
			Args = append(Args[:i], Args[i+2:]...)
			return intValue
		}
	}
	return -4541
}

var userDefined = make(map[string]string)

func UserDefinedFlags() map[string]string {
	tmp := []string{}

	tmp = append(tmp, Args...)

	for _, cmd := range tmp {
		if len(cmd) > 1 && strings.Count(cmd, "-") == 1 && strings.HasPrefix(cmd, "-") {
			value := S(cmd, cmd)
			cmd = strings.Replace(cmd, "-", "", 1)
			userDefined[cmd] = value
		}
	}

	return userDefined
}

// Read square brackets
func ReadSqBr(cmd string) (string, string) {
	c := strings.SplitN(cmd, "[", 2)
	name := c[0]
	values := c[1][:len(c[1])-1]
	return name, values
}

// Read square brackets and separate values by a string/char
func ReadSqBrSepBy(cmd string, sep string) (string, []string) {
	c := strings.SplitN(cmd, "[", 2)

	name := c[0]
	values := strings.Split(c[1][:len(c[1])-1], sep)
	return name, values
}

// Read circular brackets
func ReadCrBr(cmd string) (string, string) {
	c := strings.SplitN(cmd, "(", 2)
	name := c[0]
	values := c[1][:len(c[1])-1]
	return name, values
}

// Read circular brackets and separate values by a string/char
func ReadCrBrSepBy(cmd string, sep string) (string, []string) {
	c := strings.SplitN(cmd, "(", 2)
	name := c[0]
	values := strings.Split(c[1][:len(c[1])-1], sep)
	return name, values
}

func Parse() {

	if len(os.Args) < 2 {
		print(Help)
	}

	if showError && len(userDefined) > 0 {
		panic(fmt.Sprintf("Undefined Flags%v", userDefined))
	}
}

func init() {
	log.SetFlags(0)
}
