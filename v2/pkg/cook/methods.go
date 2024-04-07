package cook

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/glitchedgitz/cook/v2/pkg/parse"
)

func (cook *COOK) ApplyMethods(vallll []string, meths []string, array *[]string) {
	tmp := []string{}
	analyseMethods := [][]string{}

	for _, g := range meths {
		if strings.Contains(g, "[") {
			name, value := parse.ReadSqBr(g)
			analyseMethods = append(analyseMethods, []string{strings.ToLower(name), value})
		} else {
			analyseMethods = append(analyseMethods, []string{strings.ToLower(g), ""})
		}
	}

	for _, v := range analyseMethods {
		f := v[0]
		value := v[1]
		if fn, exists := cook.Method.MethodFuncs[f]; exists {
			fn(vallll, value, &tmp)
		} else if fn, exists := cook.Method.UrlFuncs[f]; exists {
			cook.Method.AnalyzeURLs(vallll, fn, &tmp)
		} else if e, exists := cook.Method.EncodersFuncs[f]; exists {
			for _, v := range vallll {
				output, err := e.Encode([]byte(v))
				if err != nil {
					log.Fatalln("Err")
				}
				tmp = append(tmp, string(output))
			}
		} else {
			fmt.Fprintf(os.Stderr, "\nMethod \"%s\" Doesn't exists\n", f)
			cook.MistypedCheck(f)
			os.Exit(0)
		}
		vallll = tmp
		tmp = nil
	}

	*array = append(*array, vallll...)
}

func (cook *COOK) MistypedCheck(mistyped string) {
	fmt.Fprintln(os.Stderr)

	fmt.Fprintln(os.Stderr, "Similar Methods")
	found := false
	check := func(k string) {
		similarity := strutil.Similarity(mistyped, k, metrics.NewHamming())
		if similarity >= 0.3 {
			fmt.Println("-", k)
			found = true
		}
	}

	for k := range cook.Method.MethodFuncs {
		check(k)
	}

	for k := range cook.Method.UrlFuncs {
		check(k)
	}

	for k := range cook.Method.EncodersFuncs {
		check(k)
	}

	if !found {
		fmt.Fprintln(os.Stderr, "None")
	}

}

func (cook *COOK) CheckMethods(p string, array *[]string) bool {
	if strings.Count(p, ".") > 0 {
		splitS := parse.SplitMethods(p)
		u := splitS[0]
		if _, exists := cook.Params[u]; exists {

			vallll := []string{}

			if !cook.CheckParam(u, &vallll) && !cook.Config.CheckYaml(u, &vallll) {
				return false
			}

			cook.ApplyMethods(vallll, splitS[1:], array)

			return true
		}

	}
	return false
}
