package cook

import (
	"fmt"
	"strings"

	"github.com/glitchedgitz/cook/v2/pkg/config"
	"github.com/glitchedgitz/cook/v2/pkg/parse"
)

func (cook *COOK) AppendMode(values []string) {
	tmp := []string{}
	till := len(cook.Final)
	if len(cook.Final) > len(values) {
		till = len(values)
	}
	for i := 0; i < till; i++ {
		tmp = append(tmp, cook.Final[i]+values[i])
	}
	cook.Final = tmp
}

func (cook *COOK) PermutationMode(values []string) {
	tmp := []string{}
	for _, t := range cook.Final {
		for _, v := range values {
			tmp = append(tmp, t+v)
		}
	}
	cook.Final = tmp
}

func (cook *COOK) CheckParam(p string, array *[]string) bool {
	if val, exists := cook.Params[p]; exists {
		if config.PipeInput(val, array) || config.RawInput(val, array) || RepeatOp(val, array) || cook.Config.ParseFunc(val, array) || cook.Config.ParseFile(p, val, array) || cook.CheckMethods(val, array) {
			return true
		}

		*array = append(*array, parse.SplitValues(val)...)
		return true
	}
	return false
}

func (cook *COOK) Print() {

	if !cook.PrintResult {
		return
	}

	if len(cook.MethodsForAll) > 0 {
		tmp := []string{}

		for _, meth := range strings.Split(cook.MethodsForAll, ",") {
			cook.ApplyMethods(cook.Final, parse.SplitMethods(meth), &tmp)
		}
		for _, v := range tmp {
			fmt.Println(v)
		}
	} else {
		for _, v := range cook.Final {
			fmt.Println(v)
		}
	}
}
