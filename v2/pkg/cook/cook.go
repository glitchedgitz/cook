package cook

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/glitchedgitz/cook/v2/pkg/config"
	"github.com/glitchedgitz/cook/v2/pkg/methods"
	"github.com/glitchedgitz/cook/v2/pkg/parse"
	"github.com/glitchedgitz/cook/v2/pkg/util"
)

func (cook *COOK) SetupConfig() {
	if cook.Config == nil {
		cook.Config = &config.Config{}
	}

	if cook.Config.HomeFolder == "" {
		cook.Config.HomeFolder, _ = os.UserHomeDir()
	}

	if cook.Config.ConfigPath == "" {
		cook.Config.ConfigPath = path.Join(cook.Config.HomeFolder, ".config", "cook")
	}

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatalln(err)
	}

	cook.Config.CachePath = path.Join(cacheDir, "cook")
	cook.Config.IngredientsPath = path.Join(cook.Config.ConfigPath, "cook-ingredients")

	cook.VPrint(fmt.Sprintln("ConfigPath: ", cook.Config.ConfigPath))
	cook.VPrint(fmt.Sprintln("IngredientsPath: ", cook.Config.IngredientsPath))
	cook.VPrint(fmt.Sprintln("CachePath: ", cook.Config.CachePath))

}

// Cook pattern can contain flags and values
// So we are parsing them here
func (cook *COOK) ParseCustomFlags() {
	parseFlags := parse.NewParse(cook.Pattern...)
	cook.Params = parseFlags.UserDefinedFlags()
	cook.Pattern = parseFlags.Args
}

func (cook *COOK) SetupMethods() {
	leetValues := make(map[string][]string)
	util.ReadInfoYaml(path.Join(cook.Config.ConfigPath, "leet.yaml"), leetValues)
	m := methods.New(leetValues)
	cook.Method = m
}

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

func NewWithoutConfig() *COOK {
	NewCook := &COOK{}
	return New(NewCook)
}

func New(newCook *COOK) *COOK {

	newCook.SetupConfig()
	newCook.ParseCustomFlags()
	newCook.SetupMethods()
	newCook.Config.CookConfig()

	newCook.TotalCols = len(newCook.Pattern)
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
