package cook

import (
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

// CookGenerator holds the shared resources and can generate multiple patterns efficiently
type CookGenerator struct {
	Config *config.Config
	Method *methods.Methods
}

// NewGenerator creates a new generator with all shared resources initialized
func NewGenerator() *CookGenerator {
	generator := &CookGenerator{
		Config: &config.Config{},
	}

	// Initialize all the shared resources
	if generator.Config.HomeFolder == "" {
		generator.Config.HomeFolder, _ = os.UserHomeDir()
	}

	if generator.Config.ConfigPath == "" {
		generator.Config.ConfigPath = path.Join(generator.Config.HomeFolder, ".config", "cook")
	}

	generator.Config.InputFile = make(map[string]bool)

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatalln(err)
	}

	generator.Config.CachePath = path.Join(cacheDir, "cook")
	generator.Config.IngredientsPath = path.Join(generator.Config.ConfigPath, "cook-ingredients")

	// Configure from file
	generator.Config.CookConfig()

	// Setup methods
	leetValues := make(map[string][]string)
	util.ReadInfoYaml(path.Join(generator.Config.ConfigPath, "leet.yaml"), leetValues)
	generator.Method = methods.New(leetValues)

	return generator
}

// Generate creates a new pattern list using the pre-initialized resources
func (g *CookGenerator) Generate(pattern []string) []string {
	// Parse flags from the pattern first
	parseFlags := parse.NewParse(pattern...)
	params := parseFlags.UserDefinedFlags()
	cleanPattern := parseFlags.Args

	// Create a new COOK instance using shared resources
	cook := &COOK{
		Config:      g.Config,
		Method:      g.Method,
		Pattern:     cleanPattern,
		Params:      params,
		PrintResult: false,
	}

	// Check for method parameters for final output
	if methodVal, ok := params["m"]; ok {
		cook.MethodsForAll = methodVal
	} else if methodVal, ok := params["method"]; ok {
		cook.MethodsForAll = methodVal
	}

	// Check for method-column parameters
	if methodColVal, ok := params["mc"]; ok {
		cook.MethodParam = methodColVal
	} else if methodColVal, ok := params["methodcol"]; ok {
		cook.MethodParam = methodColVal
	}

	// Check for append parameters
	if appendVal, ok := params["a"]; ok {
		cook.AppendParam = appendVal
	} else if appendVal, ok := params["append"]; ok {
		cook.AppendParam = appendVal
	}

	// Check for min parameter
	if minVal, ok := params["min"]; ok {
		min, err := strconv.Atoi(minVal)
		if err == nil {
			cook.Min = min
		}
	}

	// Continue with pattern-specific initialization
	cook.TotalCols = len(cook.Pattern)
	if cook.Min < 0 {
		cook.Min = cook.TotalCols
	}

	cook.SetMin()
	cook.AnalyseParams()
	cook.Final = []string{""}

	cook.AppendMap = make(map[int]bool)
	cook.MethodMap = make(map[int][]string)

	if len(cook.AppendParam) > 0 {
		cook.ParseAppend()
	}

	if len(cook.MethodParam) > 0 {
		cook.ParseMethod()
	}

	// Generate the pattern
	cook.Generate()

	// Apply methods to the final output if specified
	if len(cook.MethodsForAll) > 0 {
		methods := strings.Split(cook.MethodsForAll, ",")
		tmp := []string{}
		cook.ApplyMethods(cook.Final, methods, &tmp)
		cook.Final = tmp
	}

	return cook.Final
}

func New(newCook *COOK) *COOK {

	newCook.SetupConfig()
	newCook.Config.CookConfig()
	newCook.ParseCustomFlags()
	newCook.SetupMethods()

	newCook.TotalCols = len(newCook.Pattern)
	if newCook.Min < 0 {
		newCook.Min = newCook.TotalCols
	}

	newCook.SetMin()
	newCook.AnalyseParams()
	newCook.Final = []string{""}

	newCook.AppendMap = make(map[int]bool)
	newCook.MethodMap = make(map[int][]string)

	if len(newCook.AppendParam) > 0 {
		newCook.ParseAppend()
	}

	if len(newCook.MethodParam) > 0 {
		newCook.ParseMethod()
	}
	return newCook
}

// Search performs a search using the pre-initialized generator resources
// It searches for the given query across all ingredients in the Cook database
// Returns a slice of matching CookIngredient items and a boolean indicating if any matches were found
func (g *CookGenerator) Search(query string) ([]CookIngredient, bool) {
	// Create a temporary COOK instance with the shared config
	tempCook := &COOK{
		Config: g.Config,
		Method: g.Method,
	}

	// Use the existing Search method from the COOK struct
	return tempCook.Search(query)
}

// ApplyMethods applies specified methods to the input strings and returns the transformed result
// This function allows direct method application without pattern generation
func (g *CookGenerator) ApplyMethods(input []string, methodNames []string) ([]string, error) {
	if len(methodNames) == 0 || len(input) == 0 {
		return input, nil
	}

	// Create a temporary COOK instance with the shared resources
	tempCook := &COOK{
		Config: g.Config,
		Method: g.Method,
	}

	// Create output slice
	output := []string{}

	// Apply the methods
	err := tempCook.ApplyMethods(input, methodNames, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}
