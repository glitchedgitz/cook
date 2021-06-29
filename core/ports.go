package core

import (
	"log"
	"strconv"
	"strings"
)

func parsePorts(ports []string, array *[]string) {

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
