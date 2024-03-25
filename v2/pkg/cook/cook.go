package cook

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/glitchedgitz/cook/v2/pkg/methods"
	"github.com/glitchedgitz/cook/v2/pkg/parse"
	"github.com/glitchedgitz/cook/v2/pkg/util"
)

func New(newCook *COOK) *COOK {

	// Cook pattern can contain flags and values
	// So we are parsing them here

	parseFlags := parse.NewParse(newCook.Pattern...)
	newCook.Params = parseFlags.UserDefinedFlags()
	newCook.Pattern = parseFlags.Args

	// fmt.Print("Params: ", newCook.Params, "\n")
	// fmt.Print("Pattern: ", newCook.Pattern, "\n")

	if newCook.Config.HomeFolder == "" {
		newCook.Config.HomeFolder, _ = os.UserHomeDir()
	}

	if newCook.Config.ConfigPath == "" {
		newCook.Config.ConfigPath = path.Join(newCook.Config.HomeFolder, ".config", "cook")
	}

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatalln(err)
	}

	newCook.Config.CachePath = path.Join(cacheDir, "cook")
	newCook.Config.IngredientsPath = path.Join(newCook.Config.ConfigPath, "cook-ingredients")

	newCook.VPrint(fmt.Sprintln("ConfigPath: ", newCook.Config.ConfigPath))
	newCook.VPrint(fmt.Sprintln("IngredientsPath: ", newCook.Config.IngredientsPath))
	newCook.VPrint(fmt.Sprintln("CachePath: ", newCook.Config.CachePath))

	leetValues := make(map[string][]string)
	util.ReadInfoYaml(path.Join(newCook.Config.ConfigPath, "leet.yaml"), leetValues)
	m := methods.New(leetValues)
	newCook.Method = m

	newCook.Config.CookConfig()

	newCook.TotalCols = len(newCook.Pattern)

	newCook.analyseParams(newCook.Params)
	// cmdsMode()

	if newCook.Min < 0 {
		newCook.Min = newCook.TotalCols
	}
	newCook.SetMin()
	newCook.Final = []string{""}
	// methods.LeetBegin()

	if len(newCook.AppendParam) > 0 {
		newCook.ParseAppend()
	}

	if len(newCook.MethodParam) > 0 {
		newCook.ParseMethod()
	}
	return newCook
}
