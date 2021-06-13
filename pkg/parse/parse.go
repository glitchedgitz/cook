package parse

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var Commands = os.Args[1:]

func Bool(flag string) bool {
	for i, cmd := range Commands {
		if cmd == flag {
			Commands = append(Commands[:i], Commands[i+1:]...)
			return true
		}
	}
	return false
}

func String(flag string) string {
	for i, cmd := range Commands {
		if cmd == flag {
			value := Commands[i+1]
			Commands = append(Commands[:i], Commands[i+2:]...)
			return value
		}
	}

	return ""
}

func Int(flag string) int {
	intValue := 0
	for i, l := range Commands {
		if l == flag {
			if Commands[i+1] == "" {
				log.Fatalf("Err: Flag %s don't have value", flag)
				// min = noOfColumns - 1
			} else {
				var err error
				intValue, err = strconv.Atoi(Commands[i+1])
				// min -= 1
				if err != nil {
					log.Fatalf("Err: Flag %s needs integer value", flag)
				}
			}
			Commands = append(Commands[:i], Commands[i+2:]...)
			return intValue
		}
	}
	return 0
}

var userDefined = make(map[string]string)

func UserDefinedFlags() map[string]string {
	tmp := []string{}

	tmp = append(tmp, Commands...)

	for _, cmd := range tmp {
		if len(cmd) > 1 && strings.HasPrefix(cmd, "-") {
			value := String(cmd)
			cmd = strings.Replace(cmd, "-", "", 1)
			userDefined[cmd] = value
		}
	}

	return userDefined
}
