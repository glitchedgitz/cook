package cook

import (
	"fmt"

	"github.com/glitchedgitz/cook/v2/pkg/config"
	"github.com/glitchedgitz/cook/v2/pkg/parse"
)

func (cook *COOK) CurrentStage() {
	fmt.Println("Config----")
	fmt.Print("\tConfigPath: ", cook.Config.ConfigPath, "\n")
	fmt.Print("\tIngredientsPath: ", cook.Config.IngredientsPath, "\n")
	fmt.Print("\tCachePath: ", cook.Config.CachePath, "\n")
	// fmt.Println("cook.Config.Ingredients: ", cook.Config.Ingredients)
	// fmt.Println("cook.checkM: ", cook.Config.Ingredients)

	fmt.Println("Vars----")
	fmt.Print("\tPattern: ", cook.Pattern, "\n")
	fmt.Print("\tParams: ", cook.Params, "\n")
	fmt.Print("\tMin: ", cook.Min, "\n")
	fmt.Print("\tMethodParam: ", cook.MethodParam, "\n")
	fmt.Print("\tMethodsForAll: ", cook.MethodsForAll, "\n")
	fmt.Print("\tAppendParam: ", cook.AppendParam, "\n")
	fmt.Print("\tMethodMap: ", cook.MethodMap, "\n")
	fmt.Print("\tAppendMap: ", cook.AppendMap, "\n")
	fmt.Print("\tFinal: ", cook.Final, "\n")
	fmt.Print("\tTotalCols: ", cook.TotalCols, "\n")
}

func (cook *COOK) Generate() {
	for columnNum, param := range cook.Pattern {

		columnValues := []string{}

		for _, p := range parse.SplitValues(param) {
			if config.RawInput(p, &columnValues) || config.ParseRanges(p, &columnValues, cook.Config.Peek) || RepeatOp(p, &columnValues) || cook.CheckMethods(p, &columnValues) || cook.CheckParam(p, &columnValues) || cook.Config.CheckYaml(p, &columnValues) {
				continue
			}
			columnValues = append(columnValues, p)
		}

		// config.VPrint(fmt.Sprintf("%-40s: %s", "Time after getting values", time.Since(start)))

		if mapval, exists := cook.MethodMap[columnNum]; exists {
			tmp := []string{}
			cook.ApplyMethods(columnValues, mapval, &tmp)
			columnValues = tmp
		}

		if !cook.AppendMap[columnNum] || columnNum == 0 {
			cook.PermutationMode(columnValues)
		} else {
			cook.AppendMode(columnValues)
		}

		// config.VPrint(fmt.Sprintf("%-40s: %s", "Time ApplyColumnCases", time.Since(start)))

		if columnNum >= cook.Min {
			cook.Print()
		}
	}
}
