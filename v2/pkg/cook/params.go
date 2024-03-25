package cook

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func (cook *COOK) SetMin() {
	if cook.Min < 0 {
		cook.Min = cook.TotalCols - 1
	} else {
		if cook.Min > cook.TotalCols {
			fmt.Println("Err: min is greator than no of columns")
			os.Exit(0)
		}
		cook.Min -= 1
	}
}

func (cook *COOK) ShowCols() {
	fmt.Fprintln(os.Stderr)
	for i, p := range cook.Pattern {
		fmt.Fprintf(os.Stderr, "Col %d: %s\n", i, p)
	}
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func (cook *COOK) ParseAppend() {
	columns := strings.Split(cook.AppendParam, ",")
	for _, colNum := range columns {
		intValue, err := strconv.Atoi(colNum)
		if err != nil {
			log.Fatalf("Err: Column Value %s in not integer", colNum)
		}
		cook.AppendMap[intValue] = true
	}
}

func (cook *COOK) ParseMethod() {
	meths := strings.Split(cook.MethodParam, ";")
	forAllCols := []string{}

	var modifiedCol = make(map[int]bool)

	for _, m := range meths {
		if strings.Contains(m, ":") {
			s := strings.SplitN(m, ":", 2)
			i, err := strconv.Atoi(s[0])
			if err != nil {
				log.Fatalf("Err: \"%s\" is not integer", s[0])
			}
			if i >= cook.TotalCols {
				log.Fatalf("Err: No Column %d", i)
			}
			cook.MethodMap[i] = strings.Split(s[1], ",")
			modifiedCol[i] = true
		} else {
			forAllCols = append(forAllCols, strings.Split(m, ",")...)
		}
	}

	for i := 0; i < cook.TotalCols; i++ {
		if !modifiedCol[i] {
			cook.MethodMap[i] = forAllCols
		}
	}
}
