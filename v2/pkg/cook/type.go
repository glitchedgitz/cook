package cook

import (
	"github.com/glitchedgitz/cook/v2/pkg/config"
	"github.com/glitchedgitz/cook/v2/pkg/methods"
)

type COOK struct {
	Config        *config.Config
	Method        *methods.Methods
	Pattern       []string // pattern but now parsed
	Params        map[string]string
	Min           int
	MethodParam   string
	MethodsForAll string
	AppendParam   string
	MethodMap     map[int][]string
	AppendMap     map[int]bool
	Final         []string
	TotalCols     int
	PrintResult   bool
}
